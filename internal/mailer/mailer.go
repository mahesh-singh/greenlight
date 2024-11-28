package mailer

import (
	"bytes"
	"embed"
	"text/template"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/go-mail/mail/v2"
)

//go:embed "templates"
var templateFS embed.FS

type Mailer struct {
	dialer    *mail.Dialer
	sesClient *ses.Client
	sender    string
}

func New(host string, port int, username, password, sender string, sesClient *ses.Client) Mailer {
	dialer := mail.NewDialer(host, port, username, password)

	dialer.Timeout = 5 * time.Second

	return Mailer{
		dialer:    dialer,
		sender:    sender,
		sesClient: sesClient,
	}
}

func (m Mailer) Send(recipient, templateFile string, data interface{}) error {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)

	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)

	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)

	if err != nil {
		return err
	}

	// input := &ses.SendEmailInput{
	// 	Destination: &types.Destination{
	// 		ToAddresses: []string{
	// 			recipient,
	// 		},
	// 	},
	// 	Message: &types.Message{
	// 		Body: &types.Body{
	// 			Html: &types.Content{
	// 				Charset: aws.String("UTF-8"),
	// 				Data:    aws.String("HTML Body... from email"),
	// 			},
	// 			Text: &types.Content{
	// 				Charset: aws.String("UTF-8"),
	// 				Data:    aws.String("Plan text body from email"),
	// 			},
	// 		},
	// 		Subject: &types.Content{
	// 			Charset: aws.String("UTF-8"),
	// 			Data:    aws.String("subject.String()"),
	// 		},
	// 	},
	// 	Source: aws.String("singh.mahesh@gmail.com"),
	// }
	// result, err := m.sesClient.SendEmail(context.TODO(), input)

	// if err != nil {
	// 	return fmt.Errorf("filed to send email via AWS SES %s", err)
	// }

	// fmt.Printf("email sent ID: %s\n", *result.MessageId)

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.SetBody("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(msg)

	if err != nil {
		return err
	}

	return nil
}
