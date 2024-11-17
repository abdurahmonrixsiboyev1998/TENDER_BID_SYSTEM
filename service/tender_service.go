package service

import (
	"context"
	"tender_bid_system/model"
	"tender_bid_system/repository"
)

type TenderService struct {
	repo *repository.TenderRepository
}

func NewTenderService(repo *repository.TenderRepository) *TenderService {
	return &TenderService{repo: repo}
}

func (s *TenderService) CreateTender(ctx context.Context, tender *model.Tender) (model.Tender, error) {
	return s.repo.CreateTender(ctx, tender)
}

func (s *TenderService) ListTenders(ctx context.Context) ([]model.Tender, error) {
	return s.repo.ListTenders(ctx)
}

func (s *TenderService) UpdateTender(ctx context.Context, tender *model.Tender) (model.Tender, error) {
	return s.repo.UpdateTender(ctx, tender)
}

func (s *TenderService) DeleteTender(ctx context.Context, tenderID int) error {
	return s.repo.DeleteTender(ctx, tenderID)
}

func (s *TenderService) GetTenderByID(ctx context.Context, tenderID int) (model.Tender, error) {
	return s.repo.GetTenderByID(ctx, tenderID)
}
