package tests

import (
	"testing"

	"notification-consumer/internal/models"
	"notification-consumer/internal/processor"
)

func BenchmarkProcessor(b *testing.B) {

	n := models.Notification{
		ID:        1,
		Channel:   "EMAIL",
		Recipient: "test@gmail.com",
		Subject:   "Benchmark",
		Message:   "Benchmark Message",
	}

	for b.Loop() {
		_ = processor.Process(
			n,
			"benchmark-message-id",
		)
	}
}
