package repository

import (
	"database/sql"

	"github.com/viniciusfal/erp/infra/model"
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

	query, err := tr.connection.Prepare("INSERT INTO transactions" +
		"(id, title, value, type, category, scheduling, payment_date) " +
		"VALUES(gen_random_uuid(), $1, $2, $3, $4, $5, $6) RETURNING id")

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

func (tr *TransactionRepository) GetProducts() ([]model.Transaction, error) {
	var transactions []model.Transaction

	query := "SELECT * FROM transactions"
	rows, err := tr.connection.Query(query)
	if err != nil {
		println(err)
		return []model.Transaction{}, err
	}

	for rows.Next() {
		var transaction model.Transaction
		err = rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.Annex,
			&transaction.Category,
			&transaction.Type,
			&transaction.Payment_date,
			&transaction.Created_at,
			&transaction.Updated_at,
		)

		if err != nil {
			println(err)
			return []model.Transaction{}, err
		}

		transactions = append(transactions, transaction)
	}

	rows.Close()

	return transactions, nil

}
