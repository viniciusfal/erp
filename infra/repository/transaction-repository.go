package repository

import (
	"database/sql"
	"fmt"
	"math"
	"mime/multipart"
	"time"

	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/services"
)

type TransactionRepository struct {
	connection *sql.DB
}

func NewTransactionRepository(connection *sql.DB) TransactionRepository {
	return TransactionRepository{
		connection: connection,
	}
}

func toUtc(t *time.Time) *time.Time {
	if t != nil {
		// Convertendo para UTC
		converted := t.UTC()
		return &converted
	}
	return nil
}

func (tr *TransactionRepository) CreateTransaction(transaction model.Transaction, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	var id string

	if transaction.Payment_date != nil {
		// Se Payment_date não for nil, converte para UTC
		transaction.Payment_date = toUtc(transaction.Payment_date)
	}

	if file != nil && fileHeader != nil {
		filepath, err := services.Savefile(file, fileHeader, "./uploads")
		if err != nil {
			return "", err
		}
		transaction.Annex = &filepath
	} else {
		transaction.Annex = nil
	}

	query, err := tr.connection.Prepare("INSERT INTO transactions" +
		"(id, title, value, type, category,  scheduling, annex, payment_date, pay, details, method, nf, account) " +
		"VALUES(gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(transaction.Title, transaction.Value, transaction.Type,
		transaction.Category, transaction.Scheduling, transaction.Annex, transaction.Payment_date, transaction.Pay, transaction.Details, transaction.Method, transaction.Nf, transaction.Account).Scan(&id)
	if err != nil {
		println(err)
		return "", err
	}

	query.Close()

	return id, nil
}

func (tr *TransactionRepository) GetTransactions() ([]model.Transaction, error) {
	if tr.connection == nil {
		return nil, fmt.Errorf("conexão com o banco de dados não foi inicializada")
	}

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
			&transaction.Details,
			&transaction.Method,
			&transaction.Nf,
			&transaction.Account,
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
		&transaction.Details,
		&transaction.Method,
		&transaction.Nf,
		&transaction.Account,
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
			&transaction.Details,
			&transaction.Method,
			&transaction.Nf,
			&transaction.Account,
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

	if transaction.Payment_date != nil {
		transaction.Payment_date = toUtc(transaction.Payment_date)
	}
	transaction.Updated_at = transaction.Updated_at.UTC()
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
			pay = $8,
			details = $9,
			method = $10,
			nf = $11,
			account = $12
		WHERE
			id = $13
			`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(transaction.Title, transaction.Value, transaction.Type, transaction.Category,
		transaction.Scheduling, transaction.Annex, transaction.Payment_date, transaction.Pay, transaction.Details, transaction.Method, transaction.Nf, transaction.Account, transaction.ID)
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

func (tr *TransactionRepository) getTransactionSummaryByDate(startDate, endDate time.Time) (float64, float64, float64, error) {
	var totalEntries, totalOutcomes float64

	// Consulta SQL para somar as entradas e saídas
	query := `
		SELECT 
			SUM(CASE WHEN type = 'entrada' THEN value ELSE 0 END) AS total_entries,
			SUM(CASE WHEN type = 'saida' THEN value ELSE 0 END) AS total_outcomes
		FROM transactions 
		WHERE payment_date BETWEEN $1 AND $2`

	// Executa a consulta e obtém as somas das entradas e saídas
	err := tr.connection.QueryRow(query, startDate, endDate).Scan(&totalEntries, &totalOutcomes)
	if err != nil {
		if err == sql.ErrNoRows {
			// Se não houver transações no intervalo de datas, retornar zeros sem erro
			return 0, 0, 0, nil
		}
		// Caso contrário, loga e retorna o erro
		fmt.Println("Erro ao calcular o total das transações: ", err)
		return 0, 0, 0, err
	}

	// Calcula o balanço total
	totalBalance := totalEntries - totalOutcomes

	// Retorna as somas de entradas, saídas e o balanço total
	return totalEntries, totalOutcomes, totalBalance, nil
}

func (tr *TransactionRepository) GetTransactionGrowthByMonth() (float64, float64, float64, error) {
	now := time.Now()

	// Início e fim do mês atual
	startCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endCurrentMonth := startCurrentMonth.AddDate(0, 1, 0).Add(-time.Second)

	// Início e fim do mês anterior
	startLastMonth := startCurrentMonth.AddDate(0, -1, 0)
	endLastMonth := startCurrentMonth.Add(-time.Second)

	// Obter balanço do mês anterior
	totalEntriesLastMonth, totalOutcomesLastMonth, balanceLastMonth, err := tr.getTransactionSummaryByDate(startLastMonth, endLastMonth)
	if err != nil {
		return 0, 0, 0, err
	}

	// Obter balanço do mês atual
	totalEntriesCurrentMonth, totalOutcomesCurrentMonth, balanceCurrentMonth, err := tr.getTransactionSummaryByDate(startCurrentMonth, endCurrentMonth)
	if err != nil {
		return 0, 0, 0, err
	}

	// Evitar divisão por zero para total de entradas e saídas
	var totalEntriesGrowth, totalOutcomesGrowth float64
	if totalEntriesLastMonth != 0 {
		totalEntriesGrowth = (totalEntriesCurrentMonth - totalEntriesLastMonth) / math.Abs(totalEntriesLastMonth) * 100
	} else {
		totalEntriesGrowth = 0 // Definir como 0% se não houve entradas no mês anterior
	}

	if totalOutcomesLastMonth != 0 {
		totalOutcomesGrowth = (totalOutcomesCurrentMonth - totalOutcomesLastMonth) / math.Abs(totalOutcomesLastMonth) * 100
	} else {
		totalOutcomesGrowth = 0 // Definir como 0% se não houve saídas no mês anterior
	}

	// Calcular a taxa de crescimento do balanço entre os dois meses
	var growthRate float64
	if balanceLastMonth != 0 {
		growthRate = (balanceCurrentMonth - balanceLastMonth) / math.Abs(balanceLastMonth) * 100
	} else {
		growthRate = 0 // Se o balanço do mês anterior for zero, retornar 0% de crescimento
	}

	// Retornar a taxa de crescimento
	return totalEntriesGrowth, totalOutcomesGrowth, growthRate, nil
}
