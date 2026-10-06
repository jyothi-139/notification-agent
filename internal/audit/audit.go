package audit

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func SaveToAuditQueue(
	client *azservicebus.Client,
	body []byte,
	messageID string,
	notificationID int,
) error {

	sender, err := client.NewSender(
		"audit-queue",
		nil,
	)
	if err != nil {
		return err
	}
	defer sender.Close(context.Background())

	msg := &azservicebus.Message{
		Body: body,
		ApplicationProperties: map[string]any{
			"messageId":      messageID,
			"notificationId": notificationID,
			"status":         "SUCCESS",
			"app":            os.Getenv("APP_NAME"),
		},
	}

	return sender.SendMessage(
		context.Background(),
		msg,
		nil,
	)
}
