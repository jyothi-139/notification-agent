package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"notification-consumer/internal/audit"
	"notification-consumer/internal/dlq"
	"notification-consumer/internal/models"
	"notification-consumer/internal/monitoring"
	"notification-consumer/internal/processor"

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

			fmt.Println("Processing Notification...")
			fmt.Println("===================================")

			err = processor.Process(
				notification,
				msg.MessageID,
			)

			if err != nil {

				fmt.Printf(
					"❌ Processing Failed: %v\n",
					err,
				)

				trackFailure(
					msg.MessageID,
					fmt.Sprintf("%d", notification.ID),
					err.Error(),
				)

				retryCount := getRetryCount(msg)

				if retryCount >= MaxRetryCount {

					fmt.Println(
						"❌ Max Retries Reached",
					)

					fmt.Println(
						"❌ Moving Message To DLQ",
					)

					dlq.MoveToDLQ(
						notification,
						err.Error(),
					)

					_ = receiver.DeadLetterMessage(
						context.Background(),
						msg,
						&azservicebus.DeadLetterOptions{
							Reason: toPtr(
								"MAX_RETRIES_EXCEEDED",
							),
							ErrorDescription: toPtr(
								err.Error(),
							),
						},
					)

					continue
				}

				incrementRetryCount(msg)

				fmt.Printf(
					"Retry Attempt %d\n",
					retryCount+1,
				)

				_ = receiver.AbandonMessage(
					context.Background(),
					msg,
					nil,
				)

				continue
			}

			// ==================================
			// ARCHIVE SUCCESSFUL MESSAGE
			// ==================================

			err = audit.SaveToAuditQueue(
				client,
				msg.Body,
				msg.MessageID,
				notification.ID,
			)

			if err != nil {

				fmt.Println(
					"❌ Audit Queue Error:",
					err,
				)

				continue
			}

			fmt.Println(
				"✅ Message Archived To Audit Queue",
			)

			// ==================================
			// COMPLETE SERVICE BUS MESSAGE
			// ==================================

			err = receiver.CompleteMessage(
				context.Background(),
				msg,
				nil,
			)

			if err != nil {

				fmt.Println(
					"❌ Complete Error:",
					err,
				)

				continue
			}

			fmt.Println(
				"✅ Message Completed Successfully",
			)

			fmt.Printf(
				"✅ Notification %d Processed Successfully\n",
				notification.ID,
			)

			fmt.Printf(
				"✅ Azure Message ID: %s\n",
				msg.MessageID,
			)

			trackSuccess(
				msg.MessageID,
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
