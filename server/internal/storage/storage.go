package storage

import (
	"database/sql"
	"fmt"
	"ural-hackaton/server/internal/config"

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
		`
		CREATE TABLE IF NOT EXISTS admins
		(
			admin_id SERIAL NOT NULL,
			email character varying(256) NOT NULL,
			fullname character varying(256) NOT NULL,
			admin_password text,
			activation_token varchar(128) NOT NULL UNIQUE,
			is_activated boolean default false,
			CONSTRAINT admins_pkey PRIMARY KEY (admin_id),
			CONSTRAINT admins_email_key UNIQUE (email)
		)	
		`,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create admins table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		`
		CREATE TABLE IF NOT EXISTS debtor
		(
			debtor_id SERIAL NOT NULL,
			debtor_inn character varying(64),
			debtor_ogrnip character varying(64),
			debtor_name character varying(128),
			debtor_category character varying(256) NOT NULL,
			debtor_snils character varying(64),
			debtor_region character varying(256) NOT NULL,
			debtor_address character varying(256) NOT NULL,
			CONSTRAINT debtor_pkey PRIMARY KEY (debtor_id),
			CONSTRAINT debtor_debtor_inn_key UNIQUE (debtor_inn),
			CONSTRAINT debtor_debtor_key UNIQUE (debtor_name),
			CONSTRAINT debtor_debtor_ogrnip_key UNIQUE (debtor_ogrnip),
			CONSTRAINT debtor_debtor_snils_key UNIQUE (debtor_snils)
		)	
		`,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create debtor table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		`
		CREATE TABLE IF NOT EXISTS debtor_message
		(
			message_id SERIAL NOT NULL,
			publishing_date timestamp without time zone NOT NULL,
			message_type character varying(128) NOT NULL,
			debtor_id integer,
			message_author character varying(128) NOT NULL,
			CONSTRAINT debtor_message_pkey PRIMARY KEY (message_id),
			CONSTRAINT debtor_message_debtor_id_fkey FOREIGN KEY (debtor_id)
				REFERENCES debtor (debtor_id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)
		`,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create debtor_message table! Error: %s", err.Error()))
	}

	_, err = storage.Db.Exec(
		`
		CREATE TABLE IF NOT EXISTS tokens (
			token_id SERIAL PRIMARY KEY,
			admin_id INT UNIQUE NOT NULL REFERENCES admins(admin_id),
			refresh_token TEXT NOT NULL
		);
		`,
	)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create tokens table! Error: %s", err.Error()))
	}
}

func Init(db *sql.DB) *Storage {
	storage := Storage{
		Db: db,
	}

	return &storage
}
