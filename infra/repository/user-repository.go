package repository

import (
	"database/sql"
	"fmt"

	"github.com/viniciusfal/erp/infra/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) CreateUser(user model.User) (string, error) {
	var id string

	queryId, err := ur.connection.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer queryId.Close()

	queryId.QueryRow(user.Email)

	query, err := ur.connection.Prepare("INSERT INTO users" +
		"(id, name, password, email, rope) " +
		"VALUES(gen_random_uuid(), $1, $2, $3, $4) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer query.Close()

	err = query.QueryRow(user.Name, user.Password, user.Email, user.Rope).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return id, nil
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {

	query := "SELECT * FROM users"
	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Password,
			&user.Email,
			&user.Rope,
		)

		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (ur *UserRepository) GetUserById(id string) (*model.User, error) {
	query, err := ur.connection.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	var user model.User

	err = query.QueryRow(id).Scan(
		&user.ID,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.Rope,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &user, nil
}
