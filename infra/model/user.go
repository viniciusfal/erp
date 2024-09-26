package model

type User struct {
	ID       string `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Rope     string `json:"rope"`
}
