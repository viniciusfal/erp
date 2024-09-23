package model

import (
	"time"
)

type Transaction struct {
	ID           string     `json:"transaction_id"`
	Title        string     `json:"title"`
	Value        float64    `json:"value"`
	Type         string     `json:"type"`
	Category     string     `json:"category"`
	Scheduling   bool       `json:"scheduling"`
	Annex        *string    `json:"annex"`
	Payment_date *time.Time `json:"payment_date"`
	Created_at   time.Time  `json:"created_at"`
	Updated_at   time.Time  `json:"updated_at"`
}

func (t Transaction) PaymentDateFormat() string {
	if t.Payment_date != nil {
		return t.Payment_date.Format("2006-01-02")
	}
	return ""
}
