package model

type Cashier struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Code     string `json:"code"`
	Active   bool   `json:"active"`
	Rope     string `json:"rope"`
}
