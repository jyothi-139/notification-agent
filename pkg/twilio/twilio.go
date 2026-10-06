package twilio

import (
	"fmt"
	"os"

	twilioSdk "github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(
	to string,
	message string,
) error {

	accountSID := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_TOKEN")
	fromNumber := os.Getenv("TWILIO_PHONE")

	if accountSID == "" {
		return fmt.Errorf("TWILIO_SID is empty")
	}

	if authToken == "" {
		return fmt.Errorf("TWILIO_TOKEN is empty")
	}

	if fromNumber == "" {
		return fmt.Errorf("TWILIO_PHONE is empty")
	}

	client := twilioSdk.NewRestClientWithParams(
		twilioSdk.ClientParams{
			Username: accountSID,
			Password: authToken,
		},
	)

	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(fromNumber)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	return nil
}
