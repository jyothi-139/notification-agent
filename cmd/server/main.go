package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"notification-consumer/internal/consumer"
	"notification-consumer/internal/logger"
	"notification-consumer/internal/monitoring"
	"notification-consumer/internal/processor"
	"notification-consumer/internal/service"
	"notification-consumer/internal/worker"

	"notification-consumer/pkg/postgres"
	"notification-consumer/pkg/servicebus"
)

func main() {

	// Load .env locally only
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}

	// Initialize Logger
	logger.Init()

	// Verify App Insights Connection String
	log.Println(
		"App Insights Loaded:",
		os.Getenv("APPINSIGHTS_CONNECTION_STRING") != "",
	)

	// Initialize Application Insights
	monitoring.Init()

	// Database Connection
	db, err := postgres.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected Successfully")

	service.SetDB(db)

	// ==============================
	// Worker Pool Configuration
	// ==============================

	workerCount := 5

	if value := os.Getenv("WORKER_COUNT"); value != "" {

		n, err := strconv.Atoi(value)

		if err == nil {
			workerCount = n
		}
	}

	worker.Jobs = make(chan worker.Job, 100)

	worker.StartWorkerPool(
		workerCount,
		processor.Process,
	)

	log.Printf(
		"Worker Pool Started With %d Workers",
		workerCount,
	)

	// ==============================
	// Service Bus Configuration
	// ==============================

	connString := os.Getenv(
		"SERVICEBUS_CONNECTION_STRING",
	)

	queue := os.Getenv(
		"SERVICEBUS_QUEUE_NAME",
	)

	fmt.Println(
		"=================================",
	)
	fmt.Println(
		"CONSUMER STARTED",
	)
	fmt.Println(
		"Connection Loaded:",
		connString != "",
	)
	fmt.Println(
		"Consumer Queue:",
		queue,
	)
	fmt.Println(
		"=================================",
	)

	// Create Receiver and Client
	receiver, client, err := servicebus.NewReceiver(
		connString,
		queue,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer servicebus.CloseClient(client)

	// Start Consumer
	consumer.StartConsumer(
		client,
		receiver,
	)
}
