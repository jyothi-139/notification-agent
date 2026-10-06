package dlq

import (
	"fmt"

	"notification-consumer/internal/models"
)

func MoveToDLQ(
	n models.Notification,
	reason string,
) {

	fmt.Println(
		"==============================",
	)

	fmt.Println(
		"MESSAGE MOVED TO DLQ",
	)

	fmt.Println(
		"Notification ID:",
		n.ID,
	)

	fmt.Println(
		"Recipient:",
		n.Recipient,
	)

	fmt.Println(
		"Reason:",
		reason,
	)

	fmt.Println(
		"==============================",
	)
}
