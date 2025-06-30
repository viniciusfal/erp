package model

import (
	"time"
)

type Safe struct {
	ID          string    `json:"id"`
	Send_date   time.Time `json:"send_date" time_format:"2006-01-02"`
	Send_amount int       `json:"send_amount"`
	Active      bool      `json:"active"`
	Code        string    `json:"code"`
	Resp        string    `json:"resp"`
	Details     string    `json:"details"`
}
