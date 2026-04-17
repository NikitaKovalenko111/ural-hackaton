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
	// Do not drop schema on startup.
	// Startup should be idempotent and preserve data in existing tables.

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
		models.USERS_ROLE_LENGTH_MIGRATION,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate users role length! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.USERS_ROLE_NORMALIZATION_MIGRATION,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't normalize users roles! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.BOOKINGS_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create bookings table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.REQUESTS_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create reauests table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.ADMINS_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create admins table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.MENTORS_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create mantors table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.MENTORS_HUB_LINK_MIGRATION,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate mentors hub relation! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.REQUESTS_MENTOR_LINK_MIGRATION,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate requests mentor relation! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.EVENT_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create events table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.EVENT_MENTOR_LINK_MIGRATION,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate events mentor relation! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		models.AUTH_TOKENS_TABLE,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create auth table! Error: %s", err.Error()))
	}
}

func Init(db *sql.DB) *Storage {
	storage := Storage{
		Db: db,
	}

	return &storage
}
