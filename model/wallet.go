package model

type Wallet struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
	User_ID string  `json:"user_id"`
}
