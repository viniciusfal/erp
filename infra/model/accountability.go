package model

import "time"

type Accountability struct {
	ID              string    `json:"id"`
	Send_date       time.Time `json:"send_date"`
	Resp_id         string    `json:"resp_id"`
	Deb             float64   `json:"deb"`
	Cred            float64   `json:"cred"`
	PIX             float64   `json:"pix"`
	Coin            float64   `json:"coin"`
	Total_of_Day    float64   `json:"total_of_day"`
	Total_atlas     *float64  `json:"total_atlas"`
	Guiche          string    `json:"guiche"`
	Vias            int       `json:"vias"`
	Ter_vias        int       `json:"ter_vias"`
	Vias_atlas      int       `json:"vias_atlas"`
	Total_sec_vias  float64   `json:"total_sec_vias"`
	Total_ter_vias  float64   `json:"total_ter_vias"`
	Details         string    `json:"details"`
	Desconto        float64   `json:"desconto"`
	Annex           []string  `json:"annex"`
	Updated_at      time.Time `json:"updated_at"`
	Created_at      time.Time `json:"created_at"`
	Resp_name       *string   `json:"resp_name"`
}
