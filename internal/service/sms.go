package service

import (
	"fmt"
	"os"

	"notification-consumer/internal/models"
	"notification-consumer/pkg/twilio"
)

func SendSMS(n models.Notification) error {

	if n.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}

	if os.Getenv("MOCK_SMS") == "true" {
		fmt.Println("===================================")
		fmt.Println("[MOCK SMS]")
		fmt.Println("To      :", n.PhoneNumber)
		fmt.Println("Message :", n.Message)
		fmt.Println("===================================")

		return nil
	}

	fmt.Println("Sending SMS To:", n.PhoneNumber)

	return twilio.SendSMS(
		n.PhoneNumber,
		n.Message,
	)
}
