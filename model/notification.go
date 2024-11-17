package model

import "time"

type Notification struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Message    string    `json:"message"`
	RelationID int       `json:"relation_id"`
	Type       string    `json:"type"`
	CreatedAd  time.Time `json:"created_at"`
}
