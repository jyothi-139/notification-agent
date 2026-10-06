package twilio

import (
	"os"
	"testing"
)

func TestMissingSID(t *testing.T) {

	os.Setenv("TWILIO_SID", "")
	os.Setenv("TWILIO_TOKEN", "token")
	os.Setenv("TWILIO_PHONE", "+919999999999")

	err := SendSMS(
		"+919999999999",
		"test",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
