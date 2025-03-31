package repository

import (
	"database/sql"
	"fmt"
	"math"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/viniciusfal/erp/infra/model"

	"github.com/viniciusfal/erp/services"
	"github.com/xuri/excelize/v2"
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

	// Converte PaymentDate para UTC, se não for nil
	if transaction.Payment_date != nil {
		transaction.Payment_date = toUtc(transaction.Payment_date)
	}

	// Salva o arquivo anexo, se houver
	if file != nil && fileHeader != nil {
		filepath, err := services.Savefile(file, fileHeader, "./uploads")
		if err != nil {
			fmt.Println("Erro ao salvar o arquivo Anexo: ", err)
			return "", err
		}
		transaction.Annex = &filepath
	} else {
		transaction.Annex = nil
	}

	// Query para inserir a transação com os novos campos
	query, err := tr.connection.Prepare(`
		INSERT INTO transactions (
			id, title, value, type, category, scheduling, annex, payment_date, pay, details, method, nf, account,
			due_date, status, installment, total_installments, supplier_id) VALUES (
			gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
		) RETURNING id
	`)
	if err != nil {
		fmt.Println("Erro ao preparar a query: ", err)
		return "", err
	}

	// Executa a query com os novos campos
	err = query.QueryRow(
		transaction.Title, transaction.Value, transaction.Type,
		transaction.Category, transaction.Scheduling, transaction.Annex, transaction.Payment_date,
		transaction.Pay, transaction.Details, transaction.Method, transaction.Nf, transaction.Account,
		transaction.DueDate, transaction.Status, transaction.Installment, transaction.TotalInstallments, transaction.SupplierID,
	).Scan(&id)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return []model.Transaction{}, err
	}
	defer rows.Close()

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
			&transaction.DueDate,
			&transaction.Status,
			&transaction.Installment,
			&transaction.TotalInstallments,
			&transaction.SupplierID,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Transaction{}, err
		}

		transactions = append(transactions, transaction)
	}

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
		&transaction.DueDate,
		&transaction.Status,
		&transaction.Installment,
		&transaction.TotalInstallments,
		&transaction.SupplierID,
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

func (tr *TransactionRepository) GetTransactionsByDate(status string, startDate time.Time, endDate time.Time) ([]*model.Transaction, error) {

	query := "SELECT * FROM transactions WHERE status = $1 AND payment_date BETWEEN $2 AND $3 AND payment_date IS NOT NULL"

	rows, err := tr.connection.Query(query, status, startDate, endDate)
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
			&transaction.DueDate,
			&transaction.Status,
			&transaction.Installment,
			&transaction.TotalInstallments,
			&transaction.SupplierID,
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

func (tr *TransactionRepository) GetTodayTransactions(status string) ([]*model.Transaction, error) {
	now := time.Now()

	return tr.GetTransactionsByDate(status, now, now)
}

func (tr *TransactionRepository) GetCurreentMonthtransactions(status string) ([]*model.Transaction, error) {
	now := time.Now()

	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)

	return tr.GetTransactionsByDate(status, firstOfMonth, lastOfMonth)

}

func (tr *TransactionRepository) GetLast7DaysTransactions(status string) ([]*model.Transaction, error) {
	now := time.Now()
	lasthirdDays := now.AddDate(0, 0, -7)

	return tr.GetTransactionsByDate(status, lasthirdDays, now)
}

func (tr *TransactionRepository) GetLast30DaysTransactions(status string) ([]*model.Transaction, error) {
	now := time.Now()
	last7Days := now.AddDate(0, 0, -30)

	return tr.GetTransactionsByDate(status, last7Days, now)
}

func (tr *TransactionRepository) GetTransactionForDueDate(startDate time.Time, endDate time.Time) ([]*model.Transaction, error) {

	query := "SELECT * FROM transactions WHERE status = $1 AND due_date BETWEEN $2 AND $3"

	rows, err := tr.connection.Query(query, "aberto", startDate, endDate)
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
			&transaction.DueDate,
			&transaction.Status,
			&transaction.Installment,
			&transaction.TotalInstallments,
			&transaction.SupplierID,
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

func (tr *TransactionRepository) GetDueTodayTransactions() ([]*model.Transaction, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	return tr.GetTransactionForDueDate(startOfDay, endOfDay)
}

func (tr *TransactionRepository) GetCurreentMonthtransactionsDueDate() ([]*model.Transaction, error) {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)

	return tr.GetTransactionForDueDate(firstOfMonth, lastOfMonth)

}

func (tr *TransactionRepository) GetLast7DaysTransactionsDueDate() ([]*model.Transaction, error) {
	now := time.Now()
	lasthirdDays := now.AddDate(0, 0, -7)

	return tr.GetTransactionForDueDate(lasthirdDays, now)

}

func (tr *TransactionRepository) GetLast30DaysTransactionsDueDate() ([]*model.Transaction, error) {
	now := time.Now()
	last7Days := now.AddDate(0, 0, -30)

	return tr.GetTransactionForDueDate(last7Days, now)
}

func (tr *TransactionRepository) SetTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	// Converte PaymentDate para UTC, se não for nil
	if transaction.Payment_date != nil {
		transaction.Payment_date = toUtc(transaction.Payment_date)
	}
	transaction.Updated_at = time.Now().UTC()

	// Query para atualizar a transação com os novos campos
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
			account = $12,
			due_date = $13,
			status = $14,
			installment = $15,
			total_installments = $16,
			supplier_id = $17
		WHERE
			id = $18
	`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Executa a query com os novos campos
	_, err = query.Exec(
		transaction.Title, transaction.Value, transaction.Type, transaction.Category,
		transaction.Scheduling, transaction.Annex, transaction.Payment_date, transaction.Pay,
		transaction.Details, transaction.Method, transaction.Nf, transaction.Account,
		transaction.DueDate, transaction.Status, transaction.Installment, transaction.TotalInstallments, transaction.SupplierID,
		transaction.ID,
	)
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

func (tr *TransactionRepository) CreateTransactionsFromExcel(file multipart.File) error {
	// Abrir arquivo Excel a partir do upload
	f, err := excelize.OpenReader(file)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo Excel: %w", err)
	}
	defer f.Close()

	// Obter todas as linhas da primeira aba
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return fmt.Errorf("erro ao ler as linhas do Excel: %w", err)
	}

	// Percorrer as linhas do Excel (pulando a primeira linha que é o cabeçalho)
	for i, row := range rows {
		if i == 0 {
			continue // Pular cabeçalho
		}

		if len(row) < 8 { // Garante que a linha tem dados suficientes
			continue
		}

		// Converter os dados
		value, isNegative := parseCurrency(row[7])
		transactionType := row[7]

		if isNegative {
			transactionType = "saida"
		} else {
			transactionType = "entrada"
		}

		paymentDate, err := time.Parse("02/01/2006", row[5])
		if err != nil {
			fmt.Printf("Erro ao converter data na linha %d: %v\n", i+1, err)
			continue
		}

		transaction := model.Transaction{
			Method:       row[0],
			Category:     row[1],
			Nf:           row[2],
			Title:        row[3],
			Details:      row[4],
			Type:         transactionType,
			Payment_date: &paymentDate,
			Account:      row[6],
			Value:        value,
			Scheduling:   false,
		}

		// Criar transação no banco
		_, err = tr.CreateTransaction(transaction, nil, nil)
		if err != nil {
			fmt.Printf("Erro ao inserir transação na linha %d: %v\n", i+1, err)
		}
	}

	fmt.Println("Importação concluída!")
	return nil
}

func parseCurrency(value string) (float64, bool) {
	isNegative := strings.Contains(value, "-") // Verifica se o valor tem o sinal de negativo
	value = strings.ReplaceAll(value, "R$", "")
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",", ".")
	value = strings.ReplaceAll(value, "-", "") // Remove o sinal para converter corretamente

	parsedValue, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return 0, false // Retorna 0 caso haja erro
	}

	if isNegative {
		parsedValue *= -1 // Mantém o valor negativo
	}
	return parsedValue, isNegative
}

func (tr *TransactionRepository) CreateInstallmentTransactions(
	totalValue float64,
	totalInstallments int,
	title string,
	details string,
	transactionType string,
	initialDueDate time.Time,
	status string,
	category string,
) ([]string, error) {
	if totalInstallments <= 0 {
		return nil, fmt.Errorf("totalInstallments must be greater than zero")
	}

	var transactionIDs []string
	installmentValue := totalValue / float64(totalInstallments)
	dueDate := initialDueDate

	for i := 1; i <= totalInstallments; i++ {
		transaction := model.Transaction{
			Title:             title,
			Details:           details,
			Value:             installmentValue,
			Type:              transactionType,
			Category:          category,
			DueDate:           &dueDate,
			Status:            &status,
			Installment:       &i,
			TotalInstallments: &totalInstallments,
			Created_at:        time.Now(),
			Updated_at:        time.Now(),
		}

		// Create the transaction in the database
		id, err := tr.CreateTransaction(transaction, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create installment transaction: %w", err)
		}

		transactionIDs = append(transactionIDs, id)
		dueDate = dueDate.AddDate(0, 1, 0) // Next installment in 1 month
	}

	return transactionIDs, nil
}

func (tr *TransactionRepository) CreateLowTransaction(
	totalValue float64,
	totalInstallments int,
	title string,
	details string,
	transactionType string,
	initialDueDate time.Time,
	payment_date time.Time,
	status string,
	category string,
	nf string,
	pay bool,
	annex string,
	method string,
	account string,
) ([]string, error) {
	if totalInstallments <= 0 {
		return nil, fmt.Errorf("totalInstallments must be greater than zero")
	}

	var transactionIDs []string
	installmentValue := totalValue / float64(totalInstallments)
	dueDate := initialDueDate

	for i := 1; i <= totalInstallments; i++ {
		transaction := model.Transaction{
			Title:             title,
			Details:           details,
			Value:             installmentValue,
			Type:              transactionType,
			Category:          category,
			DueDate:           &dueDate,
			Payment_date:      &payment_date,
			Status:            &status,
			Installment:       &i,
			TotalInstallments: &totalInstallments,
			Nf:                nf,
			Annex:             &annex,
			Pay:               pay,
			Method:            method,
			Account:           account,
			Created_at:        time.Now(),
			Updated_at:        time.Now(),
		}

		// Create the transaction in the database
		id, err := tr.CreateTransaction(transaction, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create installment transaction: %w", err)
		}

		transactionIDs = append(transactionIDs, id)
		dueDate = dueDate.AddDate(0, 1, 0) // Next installment in 1 month
	}

	return transactionIDs, nil
}

func (tr *TransactionRepository) MarkLowTransaction(id string, newStatus string, payment_date *time.Time) (*model.Transaction, error) {

	// old Transaction
	transaction, err := tr.GetTransactionById(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query, err := tr.connection.Prepare(`UPDATE transactions SET status = $1, payment_date = $2 WHERE id = $3 RETURNING id, status, payment_date`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(newStatus, payment_date, id).Scan(
		&transaction.ID,
		&transaction.Status,
		&transaction.Payment_date,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return transaction, nil
}
