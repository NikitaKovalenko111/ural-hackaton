package models

const (
	HUB_TABLE = `
		CREATE TABLE users (
		user_id SERIAL PRIMARY KEY,
		user_fullname VARCHAR(50) UNIQUE NOT NULL,
		user_role VARCHAR(6) NOT NULL
		hub_id INT

		FOREIGN KEY (hub_id) REFERENCES hubs(hub_id)
	) ;
		`

	USERS_TABLE = `
		CREATE TABLE users (
		user_id SERIAL PRIMARY KEY,
		user_fullname VARCHAR(50) UNIQUE NOT NULL,
		user_role VARCHAR(6) NOT NULL
	);	
		`

	REQUESTS_TABLE = `
		CREATE TABLE requests (
		request_id SERIAL PRIMARY KEY,
		request_message TEXT NOT NULL,HU
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
		`

	ADMINS_TABLE = `
		CREATE TABLE admins (
		admin_id SERIAL PRIMARY KEY,
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
		`

	MENTORS_TABLE = `
		CREATE TABLE mentors (
		mentor_id SERIAL PRIMARY KEY,
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
		`
)
