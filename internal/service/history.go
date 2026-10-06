package service

import (
	"database/sql"
	"log"
	"strconv"

	"notification-consumer/internal/models"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

func SaveNotification(n models.Notification) {
	if DB == nil {
		log.Println("Database not initialized")
		return
	}

	_, err := DB.Exec(
		`
		INSERT INTO notifications
		(
			channel,
			recipient,
			subject,
			message,
			status
		)
		VALUES ($1,$2,$3,$4,$5)
		`,
		n.Channel,
		n.Recipient,
		n.Subject,
		n.Message,
		"PENDING",
	)

	if err != nil {
		log.Println("Notification Insert Error:", err)
		return
	}

	log.Println("Notification Saved")
}

func UpdateNotificationStatus(
	recipient string,
	status string,
) {
	if DB == nil {
		return
	}

	_, err := DB.Exec(
		`
		UPDATE notifications
		SET
			status = $1,
			updated_at = NOW()
		WHERE recipient = $2
		`,
		status,
		recipient,
	)

	if err != nil {
		log.Println("Notification Update Error:", err)
		return
	}
}

func SaveHistory(
	n models.Notification,
	status string,
	messageID string,
) {
	if DB == nil {
		log.Println("Database not initialized")
		return
	}

	_, err := DB.Exec(
		`
		INSERT INTO notification_history
		(
			notification_id,
			servicebus_message_id,
			channel,
			recipient,
			status
		)
		VALUES ($1,$2,$3,$4,$5)
		`,
		strconv.Itoa(n.ID),
		messageID,
		n.Channel,
		n.Recipient,
		status,
	)

	if err != nil {
		log.Println("History Insert Error:", err)
		return
	}

	log.Println("History Saved In DB")
}

func SaveDeliveryStatus(
	notificationID string,
	status string,
	remarks string,
) {
	if DB == nil {
		log.Println("Database not initialized")
		return
	}

	_, err := DB.Exec(
		`
		INSERT INTO delivery_status
		(
			notification_id,
			status,
			remarks
		)
		VALUES ($1,$2,$3)
		`,
		notificationID,
		status,
		remarks,
	)

	if err != nil {
		log.Println("Delivery Status Insert Error:", err)
		return
	}

	log.Println("Delivery Status Saved")
}

func SaveRetryLog(
	notificationID string,
	errorMessage string,
	retryCount int,
) {
	if DB == nil {
		log.Println("Database not initialized")
		return
	}

	_, err := DB.Exec(
		`
		INSERT INTO retry_logs
		(
			notification_id,
			error_message,
			retry_count
		)
		VALUES ($1,$2,$3)
		`,
		notificationID,
		errorMessage,
		retryCount,
	)

	if err != nil {
		log.Println("Retry Log Insert Error:", err)
		return
	}

	log.Println("Retry Log Saved")
}
