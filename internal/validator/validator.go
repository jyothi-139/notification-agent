package validator

import (
	"errors"

	"notification-consumer/internal/models"
)

func Validate(n models.Notification) error {

	switch n.Channel {

	case "EMAIL":

		if n.Recipient == "" {
			return errors.New("recipient required")
		}

		if n.Subject == "" {
			return errors.New("subject required")
		}

	case "SMS":

		if n.PhoneNumber == "" {
			return errors.New("phone number required")
		}

	case "BOTH":

		if n.Recipient == "" {
			return errors.New("recipient required")
		}

		if n.PhoneNumber == "" {
			return errors.New("phone number required")
		}

		if n.Subject == "" {
			return errors.New("subject required")
		}

	default:
		return errors.New("invalid channel")
	}

	if n.Message == "" {
		return errors.New("message required")
	}

	return nil
}
