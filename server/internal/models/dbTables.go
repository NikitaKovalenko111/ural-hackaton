package models

const (
	HUB_TABLE = `
		CREATE TABLE IF NOT EXISTS hubs
		(
			hub_id SERIAL PRIMARY KEY,
			hub_name character varying(64),
			address character varying(64),
			status character varying(64),
		)	
		`

	USERS_TABLE = `
		CREATE TABLE IF NOT EXISTS users
		(
			user_id SERIAL PRIMARY KEY,
			fullname character varying(64),
			role character varying(64),
		)	
		`

	REQUESTS_TABLE = `
		CREATE TABLE IF NOT EXISTS requests
		(
			request_id SERIAL PRIMARY KEY,
			request_message text,
			adddress character varying(64),
		)	
		`
)
