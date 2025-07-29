package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/viniciusfal/erp/infra/model"
)

type SalePointRepository struct {
	connection *sql.DB
}

func NewSalePointRepository(connection *sql.DB) *SalePointRepository {
	return &SalePointRepository{
		connection: connection,
	}
}

func (rs *SalePointRepository) Create(salePoint *model.SalePoint) (*string, error) {
	var id *string
	new_uuid := uuid.NewString()

	query, err := rs.connection.Prepare("INSERT INTO sale_points (id, guiche_name) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Printf("erro ao preparar a criação de um novo ponto de Vendas: %s", err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(new_uuid, salePoint.GuicheName).Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (rs *SalePointRepository) GetByName(name string) (*model.SalePoint, error) {
	query, err := rs.connection.Prepare("SELECT id, guiche_name FROM sale_points WHERE guiche_name = $1")
	if err != nil {
		fmt.Printf("erro ao preparar a busca de um ponto de vendas por nome: %s", err)
		return nil, err
	}
	defer query.Close()

	var salePoint model.SalePoint
	err = query.QueryRow(name).Scan(&salePoint.ID, &salePoint.GuicheName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Printf("erro ao buscar um ponto de vendas por nome: %s", err)
		return nil, err
	}

	return &salePoint, nil
}

func (rs *SalePointRepository) GetAll() ([]model.SalePoint, error) {
	query, err := rs.connection.Prepare("SELECT id, guiche_name FROM sale_points")
	if err != nil {
		fmt.Printf("erro ao preparar a busca de todos os pontos de vendas: %s", err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		fmt.Printf("erro ao buscar todos os pontos de vendas: %s", err)
		return nil, err
	}
	defer rows.Close()

	var salePoints []model.SalePoint
	for rows.Next() {
		var salePoint model.SalePoint
		err = rows.Scan(
			&salePoint.ID,
			&salePoint.GuicheName,
		)
		if err != nil {
			fmt.Printf("erro ao ler os dados do ponto de vendas: %s", err)
			return nil, err
		}
		salePoints = append(salePoints, salePoint)
	}

	return salePoints, nil
}

func (rs *SalePointRepository) Update(salePoint *model.SalePoint) error {
	query, err := rs.connection.Prepare("UPDATE sale_points SET guiche_name = $1 WHERE id = $2")
	if err != nil {
		fmt.Printf("erro ao preparar a atualização do ponto de vendas: %s", err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(salePoint.GuicheName, salePoint.ID)
	if err != nil {
		fmt.Printf("erro ao atualizar o ponto de vendas: %s", err)
		return err
	}

	return nil
}
