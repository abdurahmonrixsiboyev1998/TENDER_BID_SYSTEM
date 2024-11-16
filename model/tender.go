package model

type Tender struct {
	ID          int    `json:"id"`
	ClientID    int    `json:"client_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Budget      string `json:"budget"`
	Status      string `json:"status"`
}
