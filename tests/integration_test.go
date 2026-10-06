package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"notification-consumer/internal/models"
	"notification-consumer/internal/processor"
)

func init() {
	_ = godotenv.Load("../.env")
}

func TestIntegration(t *testing.T) {

	// Use mock SMS instead of Twilio
	os.Setenv("MOCK_SMS", "true")
	defer os.Unsetenv("MOCK_SMS")

	n := models.Notification{
		ID:          1001,
		Channel:     "BOTH",
		Recipient:   "evvsjyothi5@gmail.com",
		PhoneNumber: "+919502429626",
		Subject:     "Integration Test",
		Message:     "Testing",
	}

	err := processor.Process(
		n,
		"test-message-id-123",
	)

	if err != nil {
		t.Fatalf(
			"Integration failed: %v",
			err,
		)
	}
}
