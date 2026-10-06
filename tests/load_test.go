package tests

import (
	"os"
	"testing"

	"notification-consumer/internal/models"
	"notification-consumer/internal/service"
)

func TestLoadMockSMS(t *testing.T) {

	os.Setenv("MOCK_SMS", "true")

	for i := 1; i <= 10; i++ {

		n := models.Notification{
			ID:          i,
			Channel:     "SMS",
			PhoneNumber: "+919502429626",
			Message:     "Load Test",
		}

		if err := service.SendSMS(n); err != nil {
			t.Fatalf("failed on message %d: %v", i, err)
		}
	}
}
