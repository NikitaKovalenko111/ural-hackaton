package users_storage

import (
	"database/sql"
	"fmt"

	user_dto "ural-hackaton/internal/dto/user"
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
)

type UserRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *user_dto.CreateUserDto) error {
	_, err := r.db.Db.Exec(
		`INSERT INTO users (fullname, user_role, email, telegram, phone) VALUES ($1, $2, $3, $4, $5)`,
		user.Fullname,
		user.Role,
		user.Email,
		user.Telegram,
		user.Phone,
	)

	if err != nil {
		return fmt.Errorf("Couldn't create user!")
	}

	return nil
}

func (r *UserRepo) GetUserById(id uint64) (*models.User, error) {
	var user models.User

	err := r.db.Db.QueryRow(
		`SELECT user_id, fullname, user_role, email, telegram, phone FROM users WHERE user_id = $1`,
		id,
	).Scan(&user.Id, &user.FullName, &user.Role, &user.Email, &user.Telegram, &user.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with this id not found!")
		}

		return nil, fmt.Errorf("Couldn't get user by id!")
	}

	return &user, nil
}

func (r *UserRepo) GetUserByFullname(fullname string) (*models.User, error) {
	var user models.User

	err := r.db.Db.QueryRow(
		`SELECT user_id, fullname, user_role, email, telegram, phone FROM users WHERE fullname = $1`,
		fullname,
	).Scan(&user.Id, &user.FullName, &user.Role, &user.Email, &user.Telegram, &user.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with this fullname not found!")
		}

		return nil, fmt.Errorf("Couldn't get user by fullname!")
	}

	return &user, nil
}

func (r *UserRepo) GetUsersByRole(role string) ([]models.User, error) {
	rows, err := r.db.Db.Query(
		`SELECT user_id, fullname, user_role, email, telegram, phone FROM users WHERE user_role = $1`,
		role,
	)
	if err != nil {
		return nil, fmt.Errorf("Couldn't get users by role!")
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.FullName, &user.Role, &user.Email, &user.Telegram, &user.Phone)
		if err != nil {
			return nil, fmt.Errorf("Couldn't parse users by role!")
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Couldn't read users by role!")
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("Users with this role not found!")
	}

	return users, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Db.QueryRow(
		`SELECT user_id, fullname, user_role, email, telegram, phone FROM users WHERE email = $1`,
		email,
	).Scan(&user.Id, &user.FullName, &user.Role, &user.Email, &user.Telegram, &user.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with this fullname not found!")
		}

		return nil, fmt.Errorf("Couldn't get user by fullname!")
	}

	return &user, nil
}

func (r *UserRepo) GetUserByTelegram(telegram string) (*models.User, error) {
	var user models.User

	err := r.db.Db.QueryRow(
		`SELECT user_id, fullname, user_role, email, telegram, phone FROM users WHERE telegram = $1`,
		telegram,
	).Scan(&user.Id, &user.FullName, &user.Role, &user.Email, &user.Telegram, &user.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with this fullname not found!")
		}

		return nil, fmt.Errorf("Couldn't get user by fullname!")
	}

	return &user, nil
}
