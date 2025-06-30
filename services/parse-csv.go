package services

import (
	"encoding/csv"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/viniciusfal/erp/infra/model"
)

func ParseCSV(file multipart.File) ([]model.Transaction, error) {
	var transactions []model.Transaction

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		value, _ := strconv.ParseFloat(record[2], 64)
		payment_date, _ := time.Parse("2006-01-02", record[5]) // Ajuste o formato conforme necessário

		transaction := model.Transaction{
			ID:           uuid.New().String(),
			Title:        record[1],
			Value:        value,
			Type:         record[3],
			Category:     record[4],
			Payment_date: &payment_date,
			Details:      record[6],
			Method:       record[7],
			Nf:           record[8],
			Account:      record[9],
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
