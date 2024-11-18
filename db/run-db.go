package db

import (
	"database/sql"
	"fmt"
)

func RunDB() *sql.DB {
	DbConnection, err := ConnectDB()
	if err != nil {
		fmt.Println("cnnected")
	}

	return DbConnection

}
