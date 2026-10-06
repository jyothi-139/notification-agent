package monitoring

import (
	"log"
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var Client appinsights.TelemetryClient

func Init() {

	connStr := os.Getenv(
		"APPINSIGHTS_CONNECTION_STRING",
	)

	if connStr == "" {

		log.Println(
			"Application Insights Connection String Not Found",
		)

		return
	}

	Client = appinsights.NewTelemetryClient(
		connStr,
	)

	log.Println(
		"Application Insights Initialized Successfully",
	)

	// Test Event
	event := appinsights.NewEventTelemetry(
		"StartupTest",
	)

	event.Properties["Application"] =
		"notification-consumer"

	Client.Track(event)

	log.Println(
		"StartupTest event sent to Application Insights",
	)
}
