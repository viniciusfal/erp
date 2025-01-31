package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Usando variáveis de ambiente para configurar a conexão com o banco de dados
	host := "localhost"
	port := "5432"         // Exemplo: "5432"
	user := "postgres"     // Exemplo: "postgres"
	password := "postgres" // Exemplo: "your_password"
	dbname := "erpdb"      // Exemplo: "your_db_name"

	// Formatar a string de conexão com PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	// Definindo o timeout para a tentativa de conexão
	db.SetConnMaxLifetime(6 * time.Second)

	// Verifica se a conexão foi bem-sucedida
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	_, err = db.Exec("SET search_path TO public;")
	if err != nil {
		log.Fatalf("Error setting search_path: %v", err)
	}
	// Log de conexão bem-sucedida
	fmt.Println("Connected to database:", dbname)

	return db, nil
}
