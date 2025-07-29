package model

type Config struct {
	ID             string      `json:"id"`
	ValueSecondVia float64     `json:"value_second_via"`
	ValueterVia    float64     `json:"value_tera_via"`
	SalePoints     []SalePoint `json:"sale_points"`
	Partners       []Partner   `json:"partners"`
	TaxaCardDeb    float64     `json:"taxa_card_deb"`
	TaxaCardCred   float64     `json:"taxa_card_cred"`
}
