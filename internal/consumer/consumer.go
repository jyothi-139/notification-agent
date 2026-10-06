package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"notification-consumer/internal/models"
	"notification-consumer/internal/monitoring"
	"notification-consumer/internal/worker"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

const MaxRetryCount = 3

func StartConsumer(
	client *azservicebus.Client,
	receiver *azservicebus.Receiver,
) {

	fmt.Println("Consumer Started...")

	for {

		messages, err := receiver.ReceiveMessages(
			context.Background(),
			10,
			nil,
		)

		if err != nil {
			fmt.Println("Receive Error:", err)
			continue
		}

		for _, msg := range messages {

			fmt.Println("===================================")
			fmt.Println("MESSAGE RECEIVED")
			fmt.Printf("Message ID      : %s\n", msg.MessageID)
			fmt.Println("===================================")

			var notification models.Notification

			err = json.Unmarshal(
				msg.Body,
				&notification,
			)

			if err != nil {

				fmt.Println(
					"❌ Invalid JSON:",
					err,
				)

				trackFailure(
					msg.MessageID,
					"INVALID_PAYLOAD",
					err.Error(),
				)

				_ = receiver.DeadLetterMessage(
					context.Background(),
					msg,
					&azservicebus.DeadLetterOptions{
						Reason:           toPtr("INVALID_PAYLOAD"),
						ErrorDescription: toPtr(err.Error()),
					},
				)

				continue
			}

			trackReceived(
				msg.MessageID,
				notification.ID,
			)

			fmt.Printf(
				"Notification: %+v\n",
				notification,
			)

			fmt.Println("===================================")

			fmt.Printf(
				"Notification ID : %d\n",
				notification.ID,
			)

			fmt.Printf(
				"Channel         : %s\n",
				notification.Channel,
			)

			fmt.Printf(
				"Email           : %s\n",
				notification.Recipient,
			)

			fmt.Printf(
				"Phone Number    : %s\n",
				notification.PhoneNumber,
			)

			fmt.Println("===================================")

			worker.Jobs <- worker.Job{
				Notification: notification,
				MessageID:    msg.MessageID,
			}

			fmt.Printf(
				"✅ Notification %d queued to Worker Pool\n",
				notification.ID,
			)

			fmt.Println("===================================")
		}
	}
}

func trackReceived(
	messageID string,
	notificationID int,
) {

	if monitoring.Client == nil {
		return
	}

	event := appinsights.NewEventTelemetry(
		"MessageReceived",
	)

	event.Properties["MessageID"] =
		messageID

	event.Properties["NotificationID"] =
		fmt.Sprintf("%d", notificationID)

	monitoring.Client.Track(event)
}

func trackSuccess(
	messageID string,
	notificationID int,
) {

	if monitoring.Client == nil {
		return
	}

	event := appinsights.NewEventTelemetry(
		"MessageProcessed",
	)

	event.Properties["MessageID"] =
		messageID

	event.Properties["NotificationID"] =
		fmt.Sprintf("%d", notificationID)

	event.Properties["Status"] =
		"SUCCESS"

	monitoring.Client.Track(event)
}

func trackFailure(
	messageID string,
	notificationID string,
	errorMsg string,
) {

	if monitoring.Client == nil {
		return
	}

	event := appinsights.NewEventTelemetry(
		"MessageFailed",
	)

	event.Properties["MessageID"] =
		messageID

	event.Properties["NotificationID"] =
		notificationID

	event.Properties["Error"] =
		errorMsg

	event.Properties["Status"] =
		"FAILED"

	monitoring.Client.Track(event)
}

func toPtr(v string) *string {
	return &v
}
