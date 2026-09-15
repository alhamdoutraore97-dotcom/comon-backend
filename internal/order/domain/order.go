package domain

type Order struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}
