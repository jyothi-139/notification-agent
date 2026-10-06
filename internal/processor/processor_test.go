package processor

import (
	"testing"

	"github.com/joho/godotenv"

	"notification-consumer/internal/models"
)

func init() {
	_ = godotenv.Load("../../.env")
}

func TestProcess(t *testing.T) {

	n := models.Notification{
		ID:        1,
		Channel:   "EMAIL",
		Recipient: "test@gmail.com",
		Subject:   "Test",
		Message:   "Test Message",
	}

	err := Process(
		n,
		"test-message-id",
	)

	if err != nil {
		t.Errorf("Processing failed: %v", err)
	}
}
