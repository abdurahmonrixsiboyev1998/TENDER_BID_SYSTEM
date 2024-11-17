package repository

import (
	"context"
	"database/sql"
	"tender_bid_system/model"

	"github.com/Masterminds/squirrel"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) CreateNotification(ctx context.Context, notification *model.Notification) (model.Notification, error) {
	query, args, err := squirrel.Insert("notifications").
		Columns("user_id", "message", "relation_id", "type").
		Values(notification.UserID, notification.Message, notification.RelationID, notification.Type).
		Suffix("RETURNING id, user_id, message, relation_id, type, created_at").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return model.Notification{}, err
	}

	var notif model.Notification
	if err := r.db.QueryRow(query, args...).Scan(&notif.ID, &notif.UserID, &notif.Message, &notif.RelationID, &notif.Type, &notif.CreatedAd); err != nil {
		return model.Notification{}, err
	}
	return notif, nil
}
