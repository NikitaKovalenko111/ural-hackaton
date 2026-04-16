package storage

import (
	"database/sql"
	"fmt"
	"ural-hackaton/internal/config"
	"ural-hackaton/internal/models"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

func Connect(cfg *config.Config) *sql.DB {
	var connString = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName,
	)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to database! Error: %s", err.Error()))
	}

	return db
}

func (storage *Storage) Prepare() {
	_, err := storage.Db.Exec(
		models.HUB_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create hubs table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.USERS_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create users table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.REQUESTS_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create reauests table! Error: %s", err.Error()))
	}

}

func Init(db *sql.DB) *Storage {
	storage := Storage{
		Db: db,
	}

	return &storage
}
