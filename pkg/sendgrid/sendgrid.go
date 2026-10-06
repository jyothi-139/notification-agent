package sendgrid

import (
	"fmt"
	"os"

	sg "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(
	to string,
	subject string,
	body string,
) error {

	apiKey := os.Getenv("SENDGRID_API_KEY")
	fromEmail := os.Getenv("SENDER_EMAIL")

	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY is empty")
	}

	if fromEmail == "" {
		return fmt.Errorf("SENDER_EMAIL is empty")
	}

	message := mail.NewSingleEmail(
		mail.NewEmail(
			"Notification Agent",
			fromEmail,
		),
		subject,
		mail.NewEmail(
			"Recipient",
			to,
		),
		body,
		body,
	)

	client := sg.NewSendClient(apiKey)

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf(
			"sendgrid error [%d]: %s",
			response.StatusCode,
			response.Body,
		)
	}

	return nil
}
