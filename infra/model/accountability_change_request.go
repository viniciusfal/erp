package model

import "time"

type AccountabilityChangeRequest struct {
	ID                       string          `json:"id"`
	OriginalAccountabilityID string          `json:"original_accountability_id"`
	RequestedBy              string          `json:"requested_by"`
	ReviewedBy               *string         `json:"reviewed_by"`
	SendDate                 time.Time       `json:"send_date"`
	NewDeb                   float64         `json:"new_deb"`
	NewCred                  float64         `json:"new_cred"`
	NewPIX                   float64         `json:"new_pix"`
	NewCoin                  float64         `json:"new_coin"`
	NewTotalOfDay            float64         `json:"new_total_of_day"`
	NewVias                  int             `json:"new_vias"`
	NewGuiche                string          `json:"new_guiche"`
	Status                   string          `json:"status"` // Pendente, Aprovada, Rejeitada
	RequestReason            string          `json:"request_reason"`
	RejectionReason          string          `json:"rejection_reason"`
	CreatedAt                time.Time       `json:"created_at"`
	ReviewedAt               time.Time       `json:"reviewed_at"`
	OldAccountability        *Accountability `json:"old_accountability"`
	NewTotalSecVias          float64         `json:"new_total_sec_vias"`
	NewTerVias               int             `json:"new_ter_vias"`
	NewTotalTerVias          float64         `json:"new_total_ter_vias"`
	NewDesconto              float64         `json:"new_desconto"`
	NewAnnex                 []string        `json:"new_annex"`
}
