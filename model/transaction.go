package model

import "time"

type Transaction struct {
	ID               string    `json:"id"`
	User_ID          string    `json:"user_id"`
	Wallet_ID        string    `json:"wallet_id"`
	Category_ID      string    `json:"category_id"`
	Amount           float64   `json:"amount"`
	Description      string    `json:"description"`
	Transaction_Date time.Time `json:"transaction_date"`
	Created_At       time.Time `json:"created_at"`
}
