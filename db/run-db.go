package db

import (
	"database/sql"
	"log"
)

func RunDB() *sql.DB {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("🚨 Erro ao conectar ao banco de dados: %v", err)
	}

	return db
}
