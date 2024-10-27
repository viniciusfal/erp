package repository

import (
	"database/sql"
	"fmt"
	"time"

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
		"(id, title, value, type, category, scheduling, payment_date, pay) " +
		"VALUES(gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(transaction.Title, transaction.Value, transaction.Type,
		transaction.Category, transaction.Scheduling, transaction.Payment_date, transaction.Pay).Scan(&id)
	if err != nil {
		println(err)
		return "", err
	}

	query.Close()

	return id, nil
}

func (tr *TransactionRepository) GetTransactions() ([]model.Transaction, error) {

	query := "SELECT * FROM transactions"
	rows, err := tr.connection.Query(query)
	if err != nil {
		println(err)
		return []model.Transaction{}, err
	}

	var transactions []model.Transaction

	for rows.Next() {
		var transaction model.Transaction

		err = rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.Value,
			&transaction.Type,
			&transaction.Category,
			&transaction.Scheduling,
			&transaction.Annex,
			&transaction.Payment_date,
			&transaction.Created_at,
			&transaction.Updated_at,
			&transaction.Pay,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Transaction{}, err
		}

		transactions = append(transactions, transaction)
	}

	rows.Close()

	return transactions, nil
}

func (tr *TransactionRepository) GetTransactionById(transaction_id string) (*model.Transaction, error) {

	query, err := tr.connection.Prepare("SELECT * FROM transactions WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var transaction model.Transaction

	err = query.QueryRow(transaction_id).Scan(
		&transaction.ID,
		&transaction.Title,
		&transaction.Value,
		&transaction.Type,
		&transaction.Category,
		&transaction.Scheduling,
		&transaction.Annex,
		&transaction.Payment_date,
		&transaction.Created_at,
		&transaction.Updated_at,
		&transaction.Pay,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	query.Close()

	return &transaction, nil
}

func (tr *TransactionRepository) GetTransactionsByDate(startDate time.Time, endDate time.Time) ([]*model.Transaction, error) {

	query := "SELECT * FROM transactions WHERE payment_date BETWEEN $1 AND $2"

	rows, err := tr.connection.Query(query, startDate, endDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var transactions []*model.Transaction

	for rows.Next() {
		var transaction model.Transaction

		err = rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.Value,
			&transaction.Type,
			&transaction.Category,
			&transaction.Scheduling,
			&transaction.Annex,
			&transaction.Payment_date,
			&transaction.Created_at,
			&transaction.Updated_at,
			&transaction.Pay,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		transactions = append(transactions, &transaction)

	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return transactions, nil
}

func (tr *TransactionRepository) SetTransaction(transaction *model.Transaction) (*model.Transaction, error) {

	// old Transaction
	_, err := tr.GetTransactionById(transaction.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Update Transaction
	query, err := tr.connection.Prepare(`
		UPDATE transactions
		SET 
			title = $1,
			value = $2,
			type = $3,
			category = $4,
			scheduling = $5,
			annex = $6,
			payment_date = $7,
			updated_at = NOW(),
			pay = $8
		WHERE
			id = $9
			`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(transaction.Title, transaction.Value, transaction.Type, transaction.Category,
		transaction.Scheduling, transaction.Annex, transaction.Payment_date, transaction.Pay, transaction.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return transaction, nil
}

func (tr *TransactionRepository) RemoveTransaction(transaction_id string) error {
	_, err := tr.GetTransactionById(transaction_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	query, err := tr.connection.Prepare("DELETE FROM transactions WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(transaction_id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (tr *TransactionRepository) MarkPayment(transaction_id string) (string, error) {
	// old Transaction
	_, err := tr.GetTransactionById(transaction_id)
	if err != nil {
		fmt.Println(err)
		return "nil", err
	}

	query, err := tr.connection.Prepare(`UPDATE transactions SET pay = $1 WHERE id = $2`)
	if err != nil {
		fmt.Println(err)
		return "nil", err
	}

	defer query.Close()

	_, err = query.Exec(true, transaction_id)
	if err != nil {
		fmt.Println(err)
		return "nil", err
	}

	return transaction_id, nil

}
