package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/viniciusfal/erp/infra/model"
)

type SafeRepository struct {
	connection *sql.DB
}

func NewSafeRepository(connection *sql.DB) SafeRepository {
	return SafeRepository{
		connection: connection,
	}
}

func (sr *SafeRepository) CreateSafe(safe model.Safe) (string, error) {
	var id string

	query, err := sr.connection.Prepare("INSERT INTO safe (id, send_date, send_amount, user_resp, active)" +
		"VALUES(gen_random_uuid(), $1, $2, $3, $4) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	defer query.Close()

	err = query.QueryRow(safe.Send_date, safe.Send_amount, safe.User_resp, safe.Active).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return id, nil
}

func (sr *SafeRepository) GetSafes() ([]model.Safe, error) {
	query := "SELECT * FROM safe"

	rows, err := sr.connection.Query(query)
	if err != nil {
		println(err)
		return []model.Safe{}, err
	}

	var safes []model.Safe

	for rows.Next() {
		var safe model.Safe

		err = rows.Scan(
			&safe.ID,
			&safe.Send_date,
			&safe.Send_amount,
			&safe.User_resp,
			&safe.Active,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Safe{}, err
		}

		safes = append(safes, safe)
	}

	rows.Close()

	return safes, nil
}

func (sr *SafeRepository) SetInativeSafe(safe model.Safe) (string, error) {
	query, err := sr.connection.Prepare("UPDATE safe SET active = $1 WHERE id = $2 ")
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	defer query.Close()

	_, err = query.Exec(safe.Active, safe.ID)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return safe.ID, nil
}

func (sr *SafeRepository) SetSafe(safe *model.Safe) (*model.Safe, error) {
	query, err := sr.connection.Prepare(`
	UPDATE safe 
	SET
		send_date = $1,
		send_amount = $2
	WHERE
		id = $3
	`)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(safe.Send_date, safe.Send_amount, safe.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return safe, nil
}

func (sr *SafeRepository) GetSafeByDate(startDate time.Time, endDate time.Time) ([]*model.Safe, error) {
	query := "SELECT * FROM safe WHERE send_date BETWEEN $1 AND $2"

	rows, err := sr.connection.Query(query, startDate, endDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var safes []*model.Safe

	for rows.Next() {
		var safe model.Safe

		err = rows.Scan(
			&safe.ID,
			&safe.Send_date,
			&safe.Send_amount,
			&safe.User_resp,
			&safe.Active,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		safes = append(safes, &safe)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return safes, nil

}
