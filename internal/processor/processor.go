package processor

import (
	"strconv"
	"strings"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"

	"notification-consumer/internal/dispatcher"
	"notification-consumer/internal/models"
	"notification-consumer/internal/monitoring"
	"notification-consumer/internal/retry"
	"notification-consumer/internal/service"
	"notification-consumer/internal/validator"
)

func Process(
	n models.Notification,
	messageID string,
) error {

	// Validate Notification
	if err := validator.Validate(n); err != nil {

		service.SaveHistory(
			n,
			"VALIDATION_FAILED",
			messageID,
		)

		service.UpdateNotificationStatus(
			n.Recipient,
			"VALIDATION_FAILED",
		)

		if monitoring.Client != nil {
			monitoring.Client.TrackException(err)
			monitoring.Client.Channel().Flush()
		}

		return err
	}

	// Execute Dispatcher With Retry
	err := retry.Execute(3, func() error {
		return dispatcher.Dispatch(n)
	})

	if err != nil {

		// Invalid Input Errors
		if strings.Contains(err.Error(), "required") ||
			strings.Contains(err.Error(), "empty") {

			service.SaveHistory(
				n,
				"VALIDATION_FAILED",
				messageID,
			)

			service.UpdateNotificationStatus(
				n.Recipient,
				"VALIDATION_FAILED",
			)

			if monitoring.Client != nil {
				monitoring.Client.TrackException(err)
				monitoring.Client.Channel().Flush()
			}

			return nil
		}

		// Provider Failure
		service.SaveHistory(
			n,
			"FAILED",
			messageID,
		)

		service.UpdateNotificationStatus(
			n.Recipient,
			"FAILED",
		)

		if monitoring.Client != nil {
			monitoring.Client.TrackException(err)
			monitoring.Client.Channel().Flush()
		}

		return err
	}

	// SUCCESS
	service.SaveHistory(
		n,
		"SUCCESS",
		messageID,
	)

	service.UpdateNotificationStatus(
		n.Recipient,
		"SUCCESS",
	)

	if monitoring.Client != nil {

		event := appinsights.NewEventTelemetry(
			"NotificationSuccess",
		)

		event.Properties["NotificationID"] =
			strconv.Itoa(n.ID)

		event.Properties["MessageID"] =
			messageID

		event.Properties["Channel"] =
			n.Channel

		event.Properties["Recipient"] =
			n.Recipient

		event.Properties["Status"] =
			"SUCCESS"

		monitoring.Client.Track(event)
		monitoring.Client.Channel().Flush()
	}

	return nil
}
