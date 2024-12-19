package main

import (
	"context"
	"database/sql"
	"expvar"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/mahesh-singh/greenlight/internal/data"
	"github.com/mahesh-singh/greenlight/internal/mailer"
)

var (
	version   string
	buildTime string
)

type appConfig struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleCons  int
		maxIdleTime  string
	}

	limiter struct {
		rps    float64
		burst  int
		enable bool
	}

	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}

	aws struct {
		awsRegion string
		accessKey string
		secretKey string
	}

	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config appConfig
	logger *slog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	var cfg appConfig

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable", "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connection")
	flag.IntVar(&cfg.db.maxIdleCons, "db-max-idle-conns", 25, "PostgreSQL max idle connection")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum request per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enable, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 2525, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "00d578cc451897", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "46b159c2cc995c", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Greenlight <no-reply@greenlight.maheshlogs.dev>", "SMTP sender")

	flag.StringVar(&cfg.aws.awsRegion, "aws-region", "us-east-1", "AWS Region")
	flag.StringVar(&cfg.aws.accessKey, "aws-access-key", "", "AWS Access Key")
	flag.StringVar(&cfg.aws.secretKey, "aws-secret-key", "", "AWS Secret Key")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(s string) error {
		cfg.cors.trustedOrigins = strings.Fields(s)
		return nil
	})

	displayVersion := flag.Bool("version", false, "Display version and exist")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)

		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)

	//sesClient, err := initAWSClients(cfg)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("database connection pool established")

	defer db.Close()

	expvar.NewString("version").Set(version)

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))

	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender, nil),
	}

	err = app.serve()

	if err != nil {
		os.Exit(1)
	}

}

func openDB(cfg appConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleCons)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}

/*
func initAWSClients(cfg appConfig) (*ses.Client, error) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.aws.accessKey,
			cfg.aws.secretKey,
			"",
		)),
		config.WithRegion(cfg.aws.awsRegion))

	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config %s", err)
	}

	return ses.NewFromConfig(awsConfig), nil
}*/
