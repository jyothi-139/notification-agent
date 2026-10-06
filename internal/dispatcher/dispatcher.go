package dispatcher

import (
	"fmt"

	"notification-consumer/internal/models"
	"notification-consumer/internal/service"
)

func Dispatch(n models.Notification) error {

	fmt.Println("---------- DISPATCHER ----------")
	fmt.Println("Channel   :", n.Channel)
	fmt.Println("Email     :", n.Recipient)
	fmt.Println("Phone     :", n.PhoneNumber)
	fmt.Println("Subject   :", n.Subject)
	fmt.Println("--------------------------------")

	switch n.Channel {

	case "EMAIL":

		if n.Recipient == "" {
			return fmt.Errorf("recipient email is required")
		}

		return service.SendEmail(n)

	case "SMS":

		if n.PhoneNumber == "" {
			return fmt.Errorf("phone number is required")
		}

		return service.SendSMS(n)

	case "BOTH":

		if n.Recipient == "" {
			return fmt.Errorf("recipient email is required")
		}

		if n.PhoneNumber == "" {
			return fmt.Errorf("phone number is required")
		}

		if err := service.SendEmail(n); err != nil {
			return err
		}

		if err := service.SendSMS(n); err != nil {
			fmt.Println("SMS Error:", err)
			return err
		}
	}

	return nil
}
