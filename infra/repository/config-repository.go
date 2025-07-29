package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/viniciusfal/erp/infra/model"
)

type ConfigRepository struct {
	connection *sql.DB
}

func NewConfigRepository(connection *sql.DB) ConfigRepository {
	return ConfigRepository{
		connection: connection,
	}
}

func (rc *ConfigRepository) Create(config *model.Config) (*model.Config, error) {
	var id *string
	new_uuid := uuid.NewString()

	query, err := rc.connection.Prepare("INSERT INTO configs (id, value_second_via, value_tera_via, sale_points, partners, taxa_card_deb, taxa_card_cred) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(
		new_uuid,
		config.ValueSecondVia,
		config.ValueterVia,
		pq.Array(config.SalePoints),
		pq.Array(config.Partners),
		config.TaxaCardDeb,
		config.TaxaCardCred,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	config.ID = *id
	return config, nil
}

func (rc *ConfigRepository) GetConfig() (*model.Config, error) {
	query := "SELECT id, value_second_via, value_tera_via, sale_points, partners, taxa_card_deb, taxa_card_cred FROM configs LIMIT 1"
	
	var config model.Config
	err := rc.connection.QueryRow(query).Scan(
		&config.ID,
		&config.ValueSecondVia,
		&config.ValueterVia,
		pq.Array(&config.SalePoints),
		pq.Array(&config.Partners),
		&config.TaxaCardDeb,
		&config.TaxaCardCred,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("configuração não encontrada")
		}
		return nil, err
	}
	
	return &config, nil
}

func (rc *ConfigRepository) UpdateConfig(config *model.Config) (*model.Config, error) {
	query, err := rc.connection.Prepare("UPDATE configs SET value_second_via = $1, value_tera_via = $2, sale_points = $3, partners = $4, taxa_card_deb = $5, taxa_card_cred = $6 WHERE id = $7")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	_, err = query.Exec(
		config.ValueSecondVia,
		config.ValueterVia,
		pq.Array(config.SalePoints),
		pq.Array(config.Partners),
		config.TaxaCardDeb,
		config.TaxaCardCred,
		config.ID,
	)
	if err != nil {
		return nil, err
	}

	return config, nil
}
