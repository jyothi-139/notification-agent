package retry

import (
	"log"
	"time"
)

func Execute(maxAttempts int, fn func() error) error {
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {

		err = fn()
		if err == nil {
			return nil
		}

		log.Printf(
			"Retry Attempt %d Failed: %v",
			attempt,
			err,
		)

		if attempt < maxAttempts {

			delay := time.Duration(1<<attempt) * time.Second

			log.Printf(
				"Waiting %v before retry...",
				delay,
			)

			time.Sleep(delay)
		}
	}

	return err
}
