package model

type Partner struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	TaxaParceiro float64 `json:"taxa_parceiro"`
	GuicheName   string  `json:"guiche_name"`
}
