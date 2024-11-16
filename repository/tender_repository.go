package repository

import (
	"context"
	"database/sql"
	"tender_bid_system/model"

	"github.com/Masterminds/squirrel"
)

type TenderRepository struct {
	db *sql.DB
}

func NewTenderRepository(db *sql.DB) *TenderRepository {
	return &TenderRepository{db: db}
}

func (r *TenderRepository) CreateTender(ctx context.Context, tender *model.Tender) (model.Tender, error) {
	query, args, err := squirrel.Insert("tenders").
		Columns("client_id", "title", "description", "deadline", "budget", "status").
		Values(tender.ClientID, tender.Title, tender.Description, tender.Deadline, tender.Budget, tender.Status).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return model.Tender{}, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Tender{}, err
	}

	return *tender, nil
}

func (r *TenderRepository) ListTenders(ctx context.Context) ([]model.Tender, error) {
	query, args, err := squirrel.Select("id", "client_id", "title", "description", "deadline", "budget", "status").
		From("tenders").
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
	var tenders []model.Tender
	for rows.Next() {
		var tender model.Tender
		err := rows.Scan(&tender.ID, &tender.ClientID, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status)
		if err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}
	return tenders, nil
}

func (r *TenderRepository) UpdateTender(ctx context.Context, tender *model.Tender) (model.Tender, error) {
	query, args, err := squirrel.Update("tenders").
		Set("client_id", tender.ClientID).
		Set("title", tender.Title).
		Set("description", tender.Description).
		Set("deadline", tender.Deadline).
		Set("budget", tender.Budget).
		Set("status", tender.Status).
		Where(squirrel.Eq{"id": tender.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return model.Tender{}, err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Tender{}, err
	}
	return *tender, nil
}

func (r *TenderRepository) DeleteTender(ctx context.Context, tenderID int) error {
	query, args, err := squirrel.Delete("tenders").
		Where(squirrel.Eq{"id": tenderID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
