# Notification Agent

## Features
- Azure Service Bus Queue
- Email Notifications using SendGrid
- SMS Notifications using Twilio
- MySQL Notification History
- Logging & Validation
- Unit, Integration and Load Tests

## Run Producer

```bash
go run cmd/producer/main.go
```

## Run Consumer

```bash
go run cmd/consumer/main.go
```

## Run Tests

```bash
go test ./... -v
```

## Run Coverage

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```