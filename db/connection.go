package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Usando variáveis de ambiente para configurar a conexão com o banco de dados
	host := "postgres.railway.internal"
	port := "5432"
	user := "postgres"
	password := "UUSGxmQxAnVmLyoeUBqPyluLqNcApEHl"
	dbname := "railway"

	// Formatar a string de conexão com PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	// Verifica se a conexão foi bem-sucedida
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	// Log de conexão bem-sucedida
	fmt.Println("Connected to database:", dbname)

	return db, nil
}
