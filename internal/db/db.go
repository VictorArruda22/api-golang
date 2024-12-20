package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	User     string
	Password string
	Host     string
	Name     string
}

func Connect(config Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false", config.User, config.Password, config.Host, config.Name)
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Erro ao abrir conexão: %v. Tentando novamente...", err)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Conexão com o banco estabelecida com sucesso!")
		return db, nil
	}

	return nil, err
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Erro ao fechar a conexão: %v", err)
	} else {
		log.Println("Conexão fechada com sucesso.")
	}
}

func CreateDBConfig() Config {
	return Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
	}
}
