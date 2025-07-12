package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/viniciusfal/erp/infra/model"
)

type PartnerRepository struct {
	connection *sql.DB
}

func NewPartnerRepository(connection *sql.DB) PartnerRepository {
	return PartnerRepository{
		connection: connection,
	}
}

func (pr *PartnerRepository) Create(partner *model.Partner) (*string, error) {
	var id *string
	new_uuid := uuid.NewString()

	query, err := pr.connection.Prepare("INSERT INTO partners (id, name, taxa_parceiro, guiche_name)" +
		"VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		fmt.Printf("erro ao preparar a criação de um novo parceiro: %s", err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(new_uuid, partner.Name, partner.TaxaParceiro, partner.GuicheName).Scan(&id)
	if err != nil {
		fmt.Printf("erro ao criar um novo parceiro: %s", err)
		return nil, err
	}

	return id, nil
}

func (pr *PartnerRepository) GetById(id string) (*model.Partner, error) {
	query, err := pr.connection.Prepare("SELECT id, name, taxa_parceiro, guiche_name FROM partners WHERE id = $1")
	if err != nil {
		fmt.Printf("erro ao preparar a busca de um parceiro por id: %s", err)
		return nil, err
	}
	defer query.Close()

	var partner model.Partner
	err = query.QueryRow(id).Scan(&partner.ID, &partner.Name, &partner.TaxaParceiro, &partner.GuicheName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Printf("erro ao buscar um parceiro por id: %s", err)
		return nil, err
	}

	return &partner, nil
}

func (pr *PartnerRepository) GetAll() ([]model.Partner, error) {
	query, err := pr.connection.Prepare("SELECT id, name, taxa_parceiro, guiche_name FROM partners")
	if err != nil {
		fmt.Printf("erro ao preparar a busca de todos os parceiros: %s", err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		fmt.Printf("erro ao buscar todos os parceiros: %s", err)
		return nil, err
	}
	defer rows.Close()

	var partners []model.Partner
	for rows.Next() {
		var partner model.Partner
		err = rows.Scan(&partner.ID, &partner.Name, &partner.TaxaParceiro, &partner.GuicheName)
		if err != nil {
			fmt.Printf("erro ao ler os dados do parceiro: %s", err)
			return nil, err
		}
		partners = append(partners, partner)
	}

	return partners, nil
}
func (pr *PartnerRepository) Update(partner *model.Partner) error {
	query, err := pr.connection.Prepare("UPDATE partners SET name = $1, taxa_parceiro = $2, guiche_name = $3 WHERE id = $4")
	if err != nil {
		fmt.Printf("erro ao preparar a atualização de um parceiro: %s", err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(partner.Name, partner.TaxaParceiro, partner.GuicheName, partner.ID)
	if err != nil {
		fmt.Printf("erro ao atualizar um parceiro: %s", err)
		return err
	}

	return nil
}
