package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mutter0815/short-url-go/configs"
	_ "github.com/lib/pq"
)

func Connect(cfg *configs.DBConfig) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных %v", err)
	}
	return db

}
