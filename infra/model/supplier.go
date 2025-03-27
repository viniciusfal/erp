package model

type Supplier struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	CPF_CNPJ     string        `json:"cpf_cnpj"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Address      string        `json:"address"`
	Active       bool          `json:"active"`
	Transactions []Transaction `json:"transactions,omitempty"`
}
