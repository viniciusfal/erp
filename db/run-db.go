package db

import "database/sql"

func RunDB() *sql.DB {
	DbConnection, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	return DbConnection

}
