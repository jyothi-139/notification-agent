package retry

import (
	"errors"
	"testing"
)

func TestExecuteSuccessFirstAttempt(t *testing.T) {
	attempts := 0

	err := Execute(3, func() error {
		attempts++
		return nil
	})

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if attempts != 1 {
		t.Fatalf("expected 1 attempt, got %d", attempts)
	}
}

func TestExecuteSuccessAfterRetry(t *testing.T) {
	attempts := 0

	err := Execute(3, func() error {
		attempts++

		if attempts < 2 {
			return errors.New("temporary error")
		}

		return nil
	})

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
}

func TestExecuteFailAfterMaxAttempts(t *testing.T) {
	attempts := 0

	err := Execute(3, func() error {
		attempts++
		return errors.New("permanent error")
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}
}

func TestExecuteSingleAttemptFailure(t *testing.T) {
	attempts := 0

	err := Execute(1, func() error {
		attempts++
		return errors.New("failure")
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if attempts != 1 {
		t.Fatalf("expected 1 attempt, got %d", attempts)
	}
}
