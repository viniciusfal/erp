package repository

import (
	"database/sql"
	"fmt"

	"github.com/viniciusfal/erp/infra/model"
)

type MetaRepository struct {
	connection *sql.DB
}

func NewMetaRepository(connection *sql.DB) MetaRepository {
	return MetaRepository{
		connection: connection,
	}
}

func (mr *MetaRepository) CreateMeta(meta model.Meta) (string, error) {
	var id string

	query, err := mr.connection.Prepare("INSERT INTO meta (id, month, metaValue) VALUES(gen_random_uuid(), $1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer query.Close()

	err = query.QueryRow(meta.Month, meta.MetaValue).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return id, nil

}

func (mr *MetaRepository) GetMetas() ([]model.Meta, error) {
	query := "SELECT * FROM meta"
	rows, err := mr.connection.Query(query)
	if err != nil {
		println(err)
		return []model.Meta{}, err
	}

	var metas []model.Meta

	for rows.Next() {
		var meta model.Meta

		err = rows.Scan(
			&meta.ID,
			&meta.Month,
			&meta.MetaValue,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Meta{}, err
		}

		metas = append(metas, meta)

	}

	rows.Close()

	return metas, nil
}

func (mr MetaRepository) GetMetaByMonth(month string) (*model.Meta, error) {
	query, err := mr.connection.Prepare("SELECT * FROM meta WHERE month = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var meta model.Meta

	err = query.QueryRow(month).Scan(
		&meta.ID,
		&meta.Month,
		&meta.MetaValue,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	query.Close()

	return &meta, nil
}

func (mr *MetaRepository) SetMeta(meta model.Meta) (string, error) {
	query, err := mr.connection.Prepare("UPDATE meta SET metaValue = $1 WHERE id = $2")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer query.Close()

	_, err = query.Exec(meta.MetaValue, meta.ID)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return meta.ID, nil
}
