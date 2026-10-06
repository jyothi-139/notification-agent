package grpc

import (
	"context"

	"notification-consumer/internal/repository"
	pb "notification-consumer/proto"
)

type NotificationService struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *NotificationService) GetStatus(
	ctx context.Context,
	req *pb.StatusRequest,
) (*pb.StatusResponse, error) {

	status, messageID, err := repository.GetLatestStatus(
		req.NotificationId,
	)

	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{
		NotificationId: req.NotificationId,
		Status:         status,
		MessageId:      messageID,
	}, nil
}

func (s *NotificationService) GetHistory(
	ctx context.Context,
	req *pb.HistoryRequest,
) (*pb.HistoryResponse, error) {

	history, err := repository.GetHistory(
		req.NotificationId,
	)

	if err != nil {
		return nil, err
	}

	return &pb.HistoryResponse{
		Records: history,
	}, nil
}
