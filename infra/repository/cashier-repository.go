package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/viniciusfal/erp/infra/model"
)

type CashierRepository struct {
	connection *sql.DB
}

func NewCashierRepository(connection *sql.DB) CashierRepository {
	return CashierRepository{
		connection: connection,
	}
}

func (cr *CashierRepository) CreateCashier(cashier *model.Cashier) (*string, error) {
	var id *string
	cashier_uuid := uuid.NewString()

	query, err := cr.connection.Prepare("INSERT INTO cashier" +
		"(id, username, code, active, rope)" +
		"VALUES($1, $2, $3, true, $4) RETURNING id, username, code, active, rope")
	if err != nil {
		fmt.Printf("erro ao preparar a criação de um novo cashier: %s", err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(cashier_uuid, cashier.Username, cashier.Code, cashier.Rope).Scan(&cashier.ID, &cashier.Username, &cashier.Code, &cashier.Active, &cashier.Rope)
	if err != nil {
		fmt.Printf("erro ao executar a criação de um novo cashier: %s", err)
		return nil, err
	}

	return id, nil
}

func (cr *CashierRepository) GetCashier() ([]model.Cashier, error) {
	var cashiers []model.Cashier

	query := "SELECT * FROM cashier"
	rows, err := cr.connection.Query(query)
	if err != nil {
		fmt.Printf("Erro ao abrir conexão para buscar todos os cashier: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cashier model.Cashier

		err = rows.Scan(
			&cashier.ID,
			&cashier.Username,
			&cashier.Code,
			&cashier.Active,
			&cashier.Rope,
		)

		if err != nil {
			fmt.Printf("Erro ao percorrer a listagem de todos os cashiers: %s", err)
			return nil, err
		}

		cashiers = append(cashiers, cashier)
	}

	return cashiers, nil
}

func (cr *CashierRepository) GetCashierById(id *string) (*model.Cashier, error) {
	var cashier model.Cashier

	query, err := cr.connection.Prepare("SELECT * FROM cashier WHERE id = $1")
	if err != nil {
		fmt.Printf("Erro ao preparar a listagem de um cashier especifico, por id: %s", err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(id).Scan(
		&cashier.ID,
		&cashier.Username,
		&cashier.Code,
		&cashier.Active,
		&cashier.Rope,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Erro no sql.ErrrnoRows: %s", err)
			return nil, err
		}
		fmt.Printf("Erro ao executar a listagem do usuario por id: %s", err)
		return nil, err
	}

	return &cashier, nil
}

func (cr *CashierRepository) GetCashierByUserName(username string) (*model.Cashier, error) {
	var cashier model.Cashier

	query, err := cr.connection.Prepare("SELECT * FROM cashier WHERE username = $1")
	if err != nil {
		fmt.Printf("Erro ao preparar a listagem de um cashier especifico, por username: %s", err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(username).Scan(
		&cashier.ID,
		&cashier.Username,
		&cashier.Code,
		&cashier.Active,
		&cashier.Rope,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Erro no sql.ErrrnoRows: %s", err)
			return nil, err
		}
		fmt.Printf("Erro ao executar a listagem do usuario por username: %s", err)
		return nil, err
	}

	return &cashier, nil
}

func (cr *CashierRepository) SetCashier(cashier model.Cashier) (*model.Cashier, error) {

	query, err := cr.connection.Prepare("UPDATE cashier SET code = $1 WHERE id = $2")
	if err != nil {
		fmt.Printf("Erro ao preparar a atualização do usuário: %s", err)
		return nil, err
	}
	defer query.Close()

	_, err = query.Exec(cashier.Code, cashier.ID)
	if err != nil {
		fmt.Printf("Erro ao Executar a atualização do usuário: %s", err)
		return nil, err
	}

	return &cashier, nil
}

func (cr *CashierRepository) AlterCashier(cashier model.Cashier) (*model.Cashier, error) {

	query, err := cr.connection.Prepare("UPDATE cashier SET active = $1 WHERE id = $2")
	if err != nil {
		fmt.Printf("Erro ao preparar a atualização do status do usuário: %s", err)
		return nil, err
	}
	defer query.Close()

	_, err = query.Exec(cashier.Active, cashier.ID)
	if err != nil {
		fmt.Printf("Erro ao Executar a atualização do status do usuário: %s", err)
		return nil, err
	}

	return &cashier, nil
}
