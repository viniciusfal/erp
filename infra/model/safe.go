package model

import "time"

type Safe struct {
	ID          string    `json:"id"`
	Send_date   time.Time `json:"send_date"`
	Send_amount float64   `json:"send_amount"`
	User_resp   string    `json:"user_resp"`
	Active      bool      `json:"active"`
}
