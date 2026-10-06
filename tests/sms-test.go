package tests

import (
	"os"
	"testing"

	"notification-consumer/internal/models"
	"notification-consumer/internal/service"
)

func TestMockSMS(t *testing.T) {

	os.Setenv("MOCK_SMS", "true")

	n := models.Notification{
		ID:          1001,
		Channel:     "SMS",
		PhoneNumber: "+919502429626",
		Message:     "Test SMS",
	}

	err := service.SendSMS(n)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
