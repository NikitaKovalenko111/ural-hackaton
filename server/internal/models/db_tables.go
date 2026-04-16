package models

const (
	HUB_TABLE = `
	CREATE TABLE IF NOT EXISTS hubs (
		hub_id SERIAL PRIMARY KEY,
		hub_name VARCHAR(50) NOT NULL UNIQUE,
		address VARCHAR(100) NOT NULL UNIQUE,
		status VARCHAR(32) NOT NULL,
		city VARCHAR(64) NOT NULL,
		description TEXT,
		schedule TEXT,
		occupancy INT
	);
		`

	BOOKINGS_TABLE = `
	CREATE TABLE IF NOT EXISTS bookings (
		booking_id SERIAL PRIMARY KEY,
		booking_date TIMESTAMPTZ NOT NULL,
		booking_zone VARCHAR(64) NOT NULL,
		booking_slots INT NOT NULL,
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
	`

	USERS_TABLE = `
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		fullname VARCHAR(50) UNIQUE NOT NULL,
		user_role VARCHAR(20) NOT NULL,
		email VARCHAR(50) NOT NULL,
		telegram VARCHAR(30) NOT NULL,
		phone VARCHAR(20) UNIQUE
	);
		`

	REQUESTS_TABLE = `
	CREATE TABLE IF NOT EXISTS requests (
		request_id SERIAL PRIMARY KEY,
		request_message TEXT NOT NULL,
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
		`

	ADMINS_TABLE = `
	CREATE TABLE IF NOT EXISTS admins (
		admin_id SERIAL PRIMARY KEY,
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
		`

	MENTORS_TABLE = `
	CREATE TABLE IF NOT EXISTS mentors (
		mentor_id SERIAL PRIMARY KEY,
		user_id INT,

		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
		`

	EVENT_TABLE = `
	CREATE TABLE IF NOT EXISTS events (
		event_id SERIAL PRIMARY KEY,
		name VARCHAR(20) UNIQUE NOT NULL,
		description VARCHAR(256) NOT NULL,
		start_time TIMESTAMPTZ NOT NULL,
		end_time TIMESTAMPTZ NOT NULL,
		hub_id INT,


		FOREIGN KEY (hub_id) REFERENCES hubs(hub_id)
	);
	`
)
