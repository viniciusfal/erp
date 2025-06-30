package repository

import (
	"database/sql"
	"fmt"

	"github.com/viniciusfal/erp/infra/model"
)

type SupplierRepository struct {
	connection *sql.DB
}

func NewSupplierRepository(connection *sql.DB) *SupplierRepository {
	return &SupplierRepository{connection: connection}
}

func (r *SupplierRepository) Create(supplier *model.Supplier) (string, error) {
	var id string

	query, err := r.connection.Prepare("INSERT INTO supplier (id, name, cpf_cnpj, email, phone, address, active) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer query.Close()

	err = query.QueryRow(supplier.ID, supplier.Name, supplier.CPF_CNPJ, supplier.Email, supplier.Phone, supplier.Address, true).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return id, nil
}

func (r *SupplierRepository) FindAllWithTransactions() ([]model.Supplier, error) {
	supplierMap := make(map[string]*model.Supplier)
	var suppliers []model.Supplier

	rows, err := r.connection.Query(`
			SELECT 
					s.id, s.name, s.cpf_cnpj, s.email, s.phone, s.address, s.active,
					t.id, t.title, t.value, t.type, t.status, t.nf, t.method, t.supplier_id
			FROM supplier s
			LEFT JOIN transactions t ON s.id = t.supplier_id
			ORDER BY s.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier model.Supplier
		var transactionID, transactionTitle, transactionType, transactionNF, transactionMethod, transactionsStatus *string
		var transactionValue *float64
		var supplierID *string

		err := rows.Scan(
			&supplier.ID,
			&supplier.Name,
			&supplier.CPF_CNPJ,
			&supplier.Email,
			&supplier.Phone,
			&supplier.Address,
			&supplier.Active,
			&transactionID,
			&transactionTitle,
			&transactionValue,
			&transactionType,
			&transactionsStatus,
			&transactionNF,
			&transactionMethod,
			&supplierID,
		)
		if err != nil {
			return nil, err
		}

		// Verifica se o supplier já está no map
		if existingSupplier, exists := supplierMap[supplier.ID]; exists {
			// Adiciona apenas a transação se existir
			if transactionID != nil {
				existingSupplier.Transactions = append(existingSupplier.Transactions, model.Transaction{
					ID:         *transactionID,
					Title:      *transactionTitle,
					Value:      *transactionValue,
					Type:       *transactionType,
					Status:     transactionsStatus,
					Nf:         *transactionNF,
					Method:     *transactionMethod,
					SupplierID: supplierID,
				})
			}
		} else {
			// Cria novo supplier no map
			supplierMap[supplier.ID] = &model.Supplier{
				ID:           supplier.ID,
				Name:         supplier.Name,
				CPF_CNPJ:     supplier.CPF_CNPJ,
				Email:        supplier.Email,
				Phone:        supplier.Phone,
				Address:      supplier.Address,
				Active:       supplier.Active,
				Transactions: []model.Transaction{},
			}

			// Adiciona transação se existir
			if transactionID != nil {
				supplierMap[supplier.ID].Transactions = append(supplierMap[supplier.ID].Transactions, model.Transaction{
					ID:         *transactionID,
					Title:      *transactionTitle,
					Value:      *transactionValue,
					Type:       *transactionType,
					Status:     transactionsStatus,
					Nf:         *transactionNF,
					Method:     *transactionMethod,
					SupplierID: supplierID,
				})
			}

			// Adiciona ao slice de suppliers
			suppliers = append(suppliers, *supplierMap[supplier.ID])
		}
	}

	return suppliers, nil
}

func (r *SupplierRepository) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier

	rows, err := r.connection.Query("SELECT * FROM supplier ORDER By name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier model.Supplier
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.CPF_CNPJ, &supplier.Email, &supplier.Phone, &supplier.Address, &supplier.Active)
		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (r *SupplierRepository) FindByID(id string) (*model.Supplier, error) {
	var supplier model.Supplier
	supplier.Transactions = []model.Transaction{} // Inicializa o slice de transações

	// Query que busca o supplier e suas transações
	rows, err := r.connection.Query(`
			SELECT 
					s.id, s.name, s.cpf_cnpj, s.email, s.phone, s.address, s.active,
					t.id, t.title, t.value, t.type, t.status, t.nf, t.method, t.supplier_id
			FROM supplier s
			LEFT JOIN transactions t ON s.id = t.supplier_id
			WHERE s.id = $1
			ORDER BY s.id`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	found := false

	for rows.Next() {
		found = true

		var transactionID, transactionTitle, transactionType, transactionNF, transactionMethod *string
		var transactionValue *float64
		var transactionStatus *string
		var supplierID *string

		// Só escaneamos os dados do supplier na primeira iteração
		if len(supplier.Transactions) == 0 {
			err = rows.Scan(
				&supplier.ID,
				&supplier.Name,
				&supplier.CPF_CNPJ,
				&supplier.Email,
				&supplier.Phone,
				&supplier.Address,
				&supplier.Active,
				&transactionID,
				&transactionTitle,
				&transactionValue,
				&transactionType,
				&transactionStatus,
				&transactionNF,
				&transactionMethod,
				&supplierID,
			)
		} else {
			// Nas iterações seguintes, só precisamos das transações
			err = rows.Scan(
				nil, nil, nil, nil, nil, nil, nil, // Campos do supplier (ignorados)
				&transactionID,
				&transactionTitle,
				&transactionValue,
				&transactionType,
				&transactionStatus,
				&transactionNF,
				&transactionMethod,
				&supplierID,
			)
		}

		if err != nil {
			return nil, err
		}

		// Adiciona a transação se existir
		if transactionID != nil {
			supplier.Transactions = append(supplier.Transactions, model.Transaction{
				ID:         *transactionID,
				Title:      *transactionTitle,
				Value:      *transactionValue,
				Type:       *transactionType,
				Status:     transactionStatus,
				Nf:         *transactionNF,
				Method:     *transactionMethod,
				SupplierID: supplierID,
			})
		}
	}

	if !found {
		return nil, nil // Não encontrado
	}

	return &supplier, nil
}

func (r *SupplierRepository) Update(supplier *model.Supplier) error {
	query, err := r.connection.Prepare("UPDATE supplier SET name = $1, cpf_cnpj=$2, email = $3, phone = $4, address = $5 WHERE id = $6")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(supplier.Name, supplier.CPF_CNPJ, supplier.Email, supplier.Phone, supplier.Address, supplier.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *SupplierRepository) Desactive(id string) error {
	query, err := r.connection.Prepare("UPDATE supplier SET active = false WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
