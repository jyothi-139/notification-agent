package consumer

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func TestToPtr(t *testing.T) {
	value := "TEST"

	ptr := toPtr(value)

	if ptr == nil {
		t.Fatal("expected pointer, got nil")
	}

	if *ptr != value {
		t.Fatalf("expected %s, got %s", value, *ptr)
	}
}

func TestGetRetryCountNilProperties(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{}

	count := getRetryCount(msg)

	if count != 0 {
		t.Fatalf("expected 0, got %d", count)
	}
}

func TestGetRetryCountPropertyMissing(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{
		ApplicationProperties: map[string]any{},
	}

	count := getRetryCount(msg)

	if count != 0 {
		t.Fatalf("expected 0, got %d", count)
	}
}

func TestGetRetryCountInvalidType(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{
		ApplicationProperties: map[string]any{
			"retry_count": "abc",
		},
	}

	count := getRetryCount(msg)

	if count != 0 {
		t.Fatalf("expected 0, got %d", count)
	}
}

func TestGetRetryCountValidValue(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{
		ApplicationProperties: map[string]any{
			"retry_count": int32(5),
		},
	}

	count := getRetryCount(msg)

	if count != 5 {
		t.Fatalf("expected 5, got %d", count)
	}
}

func TestIncrementRetryCountNilProperties(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{}

	incrementRetryCount(msg)

	count := getRetryCount(msg)

	if count != 1 {
		t.Fatalf("expected 1, got %d", count)
	}
}

func TestIncrementRetryCountExistingValue(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{
		ApplicationProperties: map[string]any{
			"retry_count": int32(2),
		},
	}

	incrementRetryCount(msg)

	count := getRetryCount(msg)

	if count != 3 {
		t.Fatalf("expected 3, got %d", count)
	}
}

func TestIncrementRetryCountMultipleTimes(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{}

	incrementRetryCount(msg)
	incrementRetryCount(msg)
	incrementRetryCount(msg)

	count := getRetryCount(msg)

	if count != 3 {
		t.Fatalf("expected 3, got %d", count)
	}
}

func TestIncrementRetryCountCreatesProperty(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{}

	incrementRetryCount(msg)

	_, exists := msg.ApplicationProperties["retry_count"]

	if !exists {
		t.Fatal("retry_count property not created")
	}
}

func TestIncrementRetryCountType(t *testing.T) {
	msg := &azservicebus.ReceivedMessage{}

	incrementRetryCount(msg)

	val := msg.ApplicationProperties["retry_count"]

	_, ok := val.(int32)

	if !ok {
		t.Fatal("retry_count is not int32")
	}
}
