package usecase

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

func TestCreateTransaction(t *testing.T) {
	DB := db.RunDB()
	TransactionRepository := repository.NewTransactionRepository(DB)
	transactionUseCase := NewTransactionUseCase(TransactionRepository)

	transaction := model.Transaction{
		ID:           uuid.NewString(),
		Title:        "venda de passagens",
		Value:        decimal.NewFromFloat(200),
		Type:         "entrada",
		Category:     "Vendas Guichê",
		Scheduling:   false,
		Annex:        nil,
		Payment_date: nil,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}

	sut, err := transactionUseCase.CreateTransaction(transaction)
	if err != nil {
		t.Fatalf("Erro ao criar transação: %v", err)
	}

	if sut.ID == "" {
		t.Fatalf("Transação não foi criada corretamente, ID vazio.")
	} else {
		t.Log("Transação criada com sucesso:", sut)
	}
}
