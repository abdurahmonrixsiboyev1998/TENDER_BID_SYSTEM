package service

import (
	"context"
	"tender_bid_system/model"
	"tender_bid_system/repository"
	"time"
)

type BidService struct {
	repo *repository.BidRepository
}

func NewBidService(repo *repository.BidRepository) *BidService {
	return &BidService{repo: repo}
}

func (s *BidService) SubmitBid(ctx context.Context, bid *model.Bid) (model.Bid, error) {
	return s.repo.SubmitBid(ctx, bid)
}

func (s *BidService) ViewBidsByTenderID(ctx context.Context, tenderID int) ([]model.Bid, error) {
	return s.repo.ViewBidsByTenderID(ctx, tenderID)
}

func (s *BidService) GetBidsByPrice(ctx context.Context, price float64, delivery_time time.Time) ([]model.Bid, error) {
	return s.repo.GetBidsByPrice(ctx, price, delivery_time)
}
