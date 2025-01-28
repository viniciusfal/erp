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
	Pay          bool       `json:"pay"`
	Details      string     `json:"details"`
	Method       string     `json:"method"`
	Nf           string     `json:"nf"`
	Account      string     `json:"account"`
}
