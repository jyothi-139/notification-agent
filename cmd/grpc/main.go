package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"

	servicegrpc "notification-consumer/internal/grpc"
	pb "notification-consumer/proto"

	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env:", err)
	}

	lis, err := net.Listen(
		"tcp",
		":50051",
	)

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterNotificationServiceServer(
		grpcServer,
		&servicegrpc.NotificationService{},
	)

	log.Println("gRPC Server Running On :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
