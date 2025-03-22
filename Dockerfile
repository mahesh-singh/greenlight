FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=/app/bin/api ./cmd/api


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/api .

EXPOSE 4000


CMD ["./api"]