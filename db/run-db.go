package db

import (
	"database/sql"
	"fmt"
)

func RunDB() *sql.DB {
	DbConnection, err := ConnectDB()
	if err != nil {
		fmt.Println("Erro ao conectar com o banco de dados:", err)
		return nil
	}

	return DbConnection

}
