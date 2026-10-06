package service

import (
	"fmt"
	"os"

	"notification-consumer/internal/models"
	"notification-consumer/pkg/sendgrid"
)

func SendEmail(n models.Notification) error {

	if n.Recipient == "" {
		return fmt.Errorf("recipient email is required")
	}

	if os.Getenv("MOCK_EMAIL") == "true" {
		fmt.Println("===================================")
		fmt.Println("[MOCK EMAIL]")
		fmt.Println("To      :", n.Recipient)
		fmt.Println("Subject :", n.Subject)
		fmt.Println("Message :", n.Message)
		fmt.Println("===================================")

		return nil
	}

	fmt.Println("Sending Email To:", n.Recipient)

	return sendgrid.SendEmail(
		n.Recipient,
		n.Subject,
		n.Message,
	)
}
