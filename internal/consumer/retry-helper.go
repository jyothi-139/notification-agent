package consumer

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func getRetryCount(
	msg *azservicebus.ReceivedMessage,
) int {

	if msg.ApplicationProperties == nil {
		return 0
	}

	val, ok := msg.ApplicationProperties["retry_count"]

	if !ok {
		return 0
	}

	count, ok := val.(int32)

	if !ok {
		return 0
	}

	return int(count)
}

func incrementRetryCount(
	msg *azservicebus.ReceivedMessage,
) {

	if msg.ApplicationProperties == nil {
		msg.ApplicationProperties =
			make(map[string]any)
	}

	count := getRetryCount(msg)

	msg.ApplicationProperties["retry_count"] =
		int32(count + 1)
}
