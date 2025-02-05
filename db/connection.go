package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var Conn = RunDB()

func ConnectDB() (*sql.DB, error) {
	host := "monorail.proxy.rlwy.net"
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("PGDATABASE") // Corrigido: agora pega PGDATABASE corretamente

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %v", err)
	}

	// Definir tempo máximo de conexão (para evitar quedas prematuras)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao testar conexão com o banco de dados: %v", err)
	}

	// Configura o search_path para "public" sem encerrar a aplicação caso falhe
	_, err = db.Exec("SET search_path TO public;")
	if err != nil {
		fmt.Printf("Aviso: Falha ao definir search_path: %v\n", err)
	}

	fmt.Println("✅ Conectado ao banco de dados:", dbname)
	return db, nil
}
