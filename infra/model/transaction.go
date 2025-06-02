package model

import (
	"time"
)

type Transaction struct {
	ID                string     `json:"transaction_id"`
	Title             string     `json:"title"`
	Value             float64    `json:"value"`
	Type              string     `json:"type"` // "saida" ou "entrada"
	Category          string     `json:"category"`
	Scheduling        bool       `json:"scheduling"`
	Annex             *string    `json:"annex"`
	Payment_date      *time.Time `json:"payment_date,omitempty"` // Data de pagamento/recebimento
	Created_at        time.Time  `json:"created_at"`
	Updated_at        time.Time  `json:"updated_at"`
	Pay               bool       `json:"pay"` // Indica se foi pago/recebido
	Details           string     `json:"details"`
	Method            string     `json:"method"`                       // Método de pagamento (ex: dinheiro, cartão)
	Nf                string     `json:"nf"`                           // Número da nota fiscal
	Account           string     `json:"account"`                      // Conta associada (ex: banco, caixa)
	DueDate           *time.Time `json:"due_date,omitempty"`           // Data de vencimento
	Status            *string    `json:"status,omitempty"`             // "aberto", "pago", "recebido"
	Installment       *int       `json:"installment,omitempty"`        // Número da parcela (1, 2, 3...)
	TotalInstallments *int       `json:"total_installments,omitempty"` // Total de parcelas
	SupplierID        *string    `json:"supplier_id,omitempty"`        // Chave estrangeira para o fornecedor
}
