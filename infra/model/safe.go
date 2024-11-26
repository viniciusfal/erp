package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Safe struct {
	ID          string          `json:"id"`
	Send_date   time.Time       `json:"send_date" time_format:"2006-01-02"`
	Send_amount decimal.Decimal `json:"send_amount"`
	Active      bool            `json:"active"`
}
