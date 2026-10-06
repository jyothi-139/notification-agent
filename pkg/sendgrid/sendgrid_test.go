package sendgrid

import (
	"os"
	"testing"
)

func TestMissingAPIKey(t *testing.T) {

	os.Setenv("SENDGRID_API_KEY", "")
	os.Setenv("SENDER_EMAIL", "test@gmail.com")

	err := SendEmail(
		"user@gmail.com",
		"Test",
		"Hello",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
