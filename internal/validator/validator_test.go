package validator

import (
	"testing"

	"notification-consumer/internal/models"
)

func TestValidatorSuccess(t *testing.T) {

	n := models.Notification{
		ID:        1,
		Channel:   "EMAIL",
		Recipient: "test@gmail.com",
		Subject:   "Test",
		Message:   "Hello",
	}

	err := Validate(n)

	if err != nil {
		t.Errorf("should be valid")
	}
}

func TestValidatorFailure(t *testing.T) {

	n := models.Notification{}

	err := Validate(n)

	if err == nil {
		t.Errorf("should fail")
	}
}
