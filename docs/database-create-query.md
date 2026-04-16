# Запрос создания БД

CREATE TABLE IF NOT EXISTS users (
	user_id SERIAL PRIMARY KEY,
	fullname VARCHAR(50) UNIQUE NOT NULL,
	user_role VARCHAR(6) NOT NULL
);

CREATE TABLE IF NOT EXISTS admins (
	admin_id SERIAL PRIMARY KEY,
	user_id INT,

	FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS mentors (
	mentor_id SERIAL PRIMARY KEY,
	user_id INT,

	FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS hubs (
	hub_id SERIAL PRIMARY KEY,
	hub_name VARCHAR(50) NOT NULL UNIQUE,
	address VARCHAR(100) NOT NULL UNIQUE,
	status VARCHAR(32) NOT NULL
);

CREATE TABLE IF NOT EXISTS requests (
	request_id SERIAL PRIMARY KEY,
	request_message TEXT NOT NULL,
	user_id INT,

	FOREIGN KEY (user_id) REFERENCES users(user_id)
);