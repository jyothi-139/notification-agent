package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "notification-consumer/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewNotificationServiceClient(
		conn,
	)

	var id string

	fmt.Print("Enter Notification ID: ")
	fmt.Scanln(&id)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	statusResp, err := client.GetStatus(
		ctx,
		&pb.StatusRequest{
			NotificationId: id,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("========== STATUS ==========")
	fmt.Println("Notification ID :", statusResp.NotificationId)
	fmt.Println("Status          :", statusResp.Status)
	fmt.Println("Message ID      :", statusResp.MessageId)

	historyResp, err := client.GetHistory(
		ctx,
		&pb.HistoryRequest{
			NotificationId: id,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("========= HISTORY =========")

	if len(historyResp.Records) == 0 {
		fmt.Println("No history found")
		return
	}

	for _, record := range historyResp.Records {
		fmt.Println(record)
	}
}
