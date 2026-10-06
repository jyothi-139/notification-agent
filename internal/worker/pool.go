package worker

import (
	"log"

	"notification-consumer/internal/models"
)

type Job struct {
	Notification models.Notification
	MessageID    string
}

var Jobs chan Job

func StartWorkerPool(
	workerCount int,
	processFunc func(models.Notification, string) error,
) {
	for i := 1; i <= workerCount; i++ {

		go func(workerID int) {

			log.Printf(
				"Worker %d started",
				workerID,
			)

			for job := range Jobs {

				log.Printf(
					"Worker %d processing notification %d",
					workerID,
					job.Notification.ID,
				)

				err := processFunc(
					job.Notification,
					job.MessageID,
				)

				if err != nil {
					log.Printf(
						"Worker %d error: %v",
						workerID,
						err,
					)
				}
			}

		}(i)
	}
}
