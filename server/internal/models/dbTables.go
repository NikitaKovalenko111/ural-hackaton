package models

const (
	HUB_TABLE = `
		CREATE TABLE hubs (
		hub_id SERIAL PRIMARY KEY,
		hub_name VARCHAR(50) NOT NULL UNIQUE,
		address VARCHAR(100) NOT NULL UNIQUE,
		status VARCHAR(32) NOT NULL
	);
		`

	USERS_TABLE = `
		CREATE TABLE users (
		user_id SERIAL PRIMARY KEY,
		fullname VARCHAR(50) UNIQUE NOT NULL,
		user_role VARCHAR(6) NOT NULL
	);	
		`

	REQUESTS_TABLE = `
		CREATE TABLE requests (
		request_id SERIAL PRIMARY KEY,
		request_message TEXT NOT NULL,
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
		`
)
