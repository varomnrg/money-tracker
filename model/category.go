package model

type Category struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	User_ID string `json:"user_id"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}
	