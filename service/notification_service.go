package service

import (
	"context"
	"tender_bid_system/model"
	"tender_bid_system/repository"
)

type NotificationService struct {
	repo *repository.NotificationRepository
}

func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) CreateNotification(ctx context.Context, notification *model.Notification) (model.Notification, error) {
	return s.repo.CreateNotification(ctx, notification)
}
