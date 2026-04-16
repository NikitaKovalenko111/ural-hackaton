package requests_storage

import (
	"database/sql"

	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type RequestRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *RequestRepo {
	return &RequestRepo{
		db: db,
	}
}

func (r *RequestRepo) CreateRequest(message string, userId uint64) error {
	_, err := r.db.Db.Exec("INSERT INTO requests (request_message, user_id) VALUES ($1, $2)", message, userId)

	return err
}

func (r *RequestRepo) GetRequestById(id uint64) (*models.Requests, error) {
	var request models.Requests

	err := r.db.Db.QueryRow(
		`SELECT request_id, request_message, user_id FROM requests WHERE request_id = $1`,
		id,
	).Scan(&request.Id, &request.Message, &request.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fiber.NewError(fiber.StatusNotFound, "Request with this id not found!")
		}

		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't get request by id!")
	}

	return &request, nil
}

func (r *RequestRepo) GetRequestsByMessage(message string) ([]models.Requests, error) {
	rows, err := r.db.Db.Query(
		`SELECT request_id, request_message, user_id FROM requests WHERE request_message = $1`,
		message,
	)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't get requests by message!")
	}
	defer rows.Close()

	requests := make([]models.Requests, 0)
	for rows.Next() {
		var request models.Requests
		err = rows.Scan(&request.Id, &request.Message, &request.UserId)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't parse requests by message!")
		}

		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't read requests by message!")
	}

	if len(requests) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Requests with this message not found!")
	}

	return requests, nil
}

func (r *RequestRepo) GetRequestsByUserId(userId uint64) ([]models.Requests, error) {
	rows, err := r.db.Db.Query(
		`SELECT request_id, request_message, user_id FROM requests WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't get requests by user id!")
	}
	defer rows.Close()

	requests := make([]models.Requests, 0)
	for rows.Next() {
		var request models.Requests
		err = rows.Scan(&request.Id, &request.Message, &request.UserId)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't parse requests by user id!")
		}

		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Couldn't read requests by user id!")
	}

	if len(requests) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Requests for this user not found!")
	}

	return requests, nil
}
