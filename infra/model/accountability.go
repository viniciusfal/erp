package model

import "time"

type Accountability struct {
	ID           string    `json:"id"`
	Send_date    time.Time `json:"send_date"`
	Resp_id      string    `json:"resp_id"`
	Deb          float64   `json:"deb"`
	PIX          float64   `json:"pix"`
	Coin         float64   `json:"coin"`
	Total_of_Day float64   `json:"total_of_day"`
	Vias         int       `json:"vias"`
	Guiche       string    `json:"guiche"`
	Created_at   time.Time `json:"created_at"`
	Resp_name    *string   `json:"resp_name"`
}
