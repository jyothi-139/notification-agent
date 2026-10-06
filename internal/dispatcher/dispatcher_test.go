package dispatcher

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"notification-consumer/internal/models"
)

func init() {
	_ = godotenv.Load("../../.env")
}

func TestDispatchEmail(t *testing.T) {
	n := models.Notification{
		ID:        1,
		Channel:   "EMAIL",
		Recipient: "test@gmail.com",
		Subject:   "Test",
		Message:   "Hello",
	}

	if err := Dispatch(n); err != nil {
		t.Fatalf("Expected nil but got %v", err)
	}
}

func TestDispatchSMS(t *testing.T) {
	os.Setenv("MOCK_SMS", "true")
	defer os.Unsetenv("MOCK_SMS")

	n := models.Notification{
		ID:          2,
		Channel:     "SMS",
		PhoneNumber: "+919502429626",
		Message:     "Hello",
	}

	if err := Dispatch(n); err != nil {
		t.Fatalf("Expected nil but got %v", err)
	}
}

func TestDispatchBoth(t *testing.T) {
	os.Setenv("MOCK_SMS", "true")
	defer os.Unsetenv("MOCK_SMS")

	n := models.Notification{
		ID:          3,
		Channel:     "BOTH",
		Recipient:   "test@gmail.com",
		PhoneNumber: "+919502429626",
		Subject:     "Test Both",
		Message:     "Hello Both",
	}

	if err := Dispatch(n); err != nil {
		t.Fatalf("Expected nil but got %v", err)
	}
}
