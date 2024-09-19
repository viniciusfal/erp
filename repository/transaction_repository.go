package repository

import (
	"database/sql"

	"github.com/viniciusfal/finances/model"
)

type TransactionRepository struct {
	connection *sql.DB
}

func NewTransactionRepository(connection *sql.DB) TransactionRepository {
	return TransactionRepository{
		connection: connection,
	}
}

func (tr *TransactionRepository) CreateTransaction(transaction model.Transaction) (string, error) {
	var id string

	query, err := tr.connection.Prepare("INSERT INTO transaction" +
		"(title, value, type, category, scheduling, payment_date) " +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING id")

	if err != nil {
		println(err)
		return "", err
	}

	err = query.QueryRow(transaction.Title, transaction.Value, transaction.Type,
		transaction.Category, transaction.Scheduling, transaction.Payment_date).Scan(&id)
	if err != nil {
		println(err)
		return "", err
	}

	query.Close()

	return id, nil

}
