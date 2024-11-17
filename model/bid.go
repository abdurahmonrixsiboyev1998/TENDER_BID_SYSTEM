package model

type Bid struct {
	ID           int     `json:"id"`
	TenderID     int     `json:"tender_id"`
	ContraktorID int     `json:"contractor_id"`
	Price        float64 `json:"price"`
	DeliveryTime int     `json:"delivery_time"`
	Comments     string  `json:"comments"`
	Status       string  `json:"status"`
}
