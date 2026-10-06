package repository

import (
	"notification-consumer/pkg/postgres"
)

func GetLatestStatus(
	notificationID string,
) (
	string,
	string,
	error,
) {

	db, err := postgres.NewConnection()
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	var status string
	var messageID string

	err = db.QueryRow(
		`
		SELECT
			status,
			COALESCE(servicebus_message_id, '')
		FROM notification_history
		WHERE notification_id = $1
		ORDER BY id DESC
		LIMIT 1
		`,
		notificationID,
	).Scan(
		&status,
		&messageID,
	)

	if err != nil {
		return "", "", err
	}

	return status, messageID, nil
}

func GetHistory(
	notificationID string,
) ([]string, error) {

	db, err := postgres.NewConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(
		`
		SELECT
			recipient,
			channel,
			status,
			COALESCE(servicebus_message_id, '')
		FROM notification_history
		WHERE notification_id = $1
		ORDER BY id DESC
		`,
		notificationID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []string

	for rows.Next() {

		var recipient string
		var channel string
		var status string
		var messageID string

		err := rows.Scan(
			&recipient,
			&channel,
			&status,
			&messageID,
		)

		if err != nil {
			return nil, err
		}

		record :=
			"Recipient: " + recipient +
				" | Channel: " + channel +
				" | Status: " + status +
				" | AzureMessageID: " + messageID

		records = append(
			records,
			record,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
