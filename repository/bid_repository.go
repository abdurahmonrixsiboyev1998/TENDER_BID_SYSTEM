package repository

import (
	"context"
	"database/sql"
	"tender_bid_system/model"
	"time"

	"github.com/Masterminds/squirrel"
)

type BidRepository struct {
	db *sql.DB
}

func NewBidRepository(db *sql.DB) *BidRepository {
	return &BidRepository{db: db}
}

func (r *BidRepository) SubmitBid(ctx context.Context, bid *model.Bid) (model.Bid, error) {
	query, args, err := squirrel.Insert("bids").
		Columns("tender_id", "contractor_id", "price", "delivery_time", "comments", "status").
		Values(bid.TenderID, bid.ContraktorID, bid.Price, bid.DeliveryTime, bid.Comments, bid.Status).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return model.Bid{}, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Bid{}, err
	}

	return *bid, nil
}

func (r *BidRepository) ViewBidsByTenderID(ctx context.Context, tenderID int) ([]model.Bid, error) {
	query, args, err := squirrel.Select("*").
		From("bids").
		Where(squirrel.Eq{"tender_id": tenderID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bids []model.Bid
	for rows.Next() {
		var bid model.Bid
		err := rows.Scan(&bid.ID, &bid.TenderID, &bid.ContraktorID, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status)
		if err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}
	return bids, nil
}

func (r *BidRepository) GetBidsByPrice(ctx context.Context, price float64, delivery_time time.Time) ([]model.Bid, error) {
	query, args, err := squirrel.Select("*").
		From("bids").
		Where(squirrel.Eq{"price": price, "delivery_time": delivery_time, "status": "pending"}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bids []model.Bid
	for rows.Next() {
		var bid model.Bid
		err := rows.Scan(&bid.ID, &bid.TenderID, &bid.ContraktorID, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status)
		if err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}
	return bids, nil
}
