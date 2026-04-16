package users_storage

import (
	"database/sql"

	usersDto "ural-hackaton/internal/dto/user"
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type UserRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *usersDto.CreateUserDto) error {
	_, err := r.db.Db.Exec(
		`INSERT INTO users (fullname, user_role) VALUES ($1, $2)`,
		user.Fullname,
		user.Role,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't create user!")
	}

	return nil
}

func (r *UserRepo) GetUserById(id uint64) (*models.User, error) {
	var user models.User

	err := r.db.Db.QueryRow(
		`SELECT user_id, user_fullname, user_role FROM users WHERE user_id = $1`,
		id,
	).Scan(&user.Id, &user.FullName, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fiber.NewError(fiber.StatusNotFound, "User with this id not found!")
		}

		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't get user by id!")
	}

	return &user, nil
}

func (r *UserRepo) GetUserByFullname(fullname string) (*models.User, error) {
	var user models.User

	err := r.db.Db.QueryRow(
		`SELECT user_id, user_fullname, user_role FROM users WHERE user_fullname = $1`,
		fullname,
	).Scan(&user.Id, &user.FullName, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fiber.NewError(fiber.StatusNotFound, "User with this fullname not found!")
		}

		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't get user by fullname!")
	}

	return &user, nil
}

func (r *UserRepo) GetUsersByRole(role string) ([]models.User, error) {
	rows, err := r.db.Db.Query(
		`SELECT user_id, user_fullname, user_role FROM users WHERE user_role = $1`,
		role,
	)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't get users by role!")
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.FullName, &user.Role)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't parse users by role!")
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't read users by role!")
	}

	if len(users) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Users with this role not found!")
	}

	return users, nil
}
