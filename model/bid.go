package model

import "time"

type Bid struct {
	ID           int       `json:"id"`
	TenderID     int       `json:"tender_id"`
	ContraktorID int       `json:"contractor_id"`
	Price        int       `json:"price"`
	DeliveryTime time.Time `json:"delivery_time"`
	Comments     string    `json:"comments"`
	Status       string    `json:"status"`
}
