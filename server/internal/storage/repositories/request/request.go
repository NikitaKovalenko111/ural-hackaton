package requests_storage

import (
	"database/sql"

	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
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
			return nil, err
		}

		return nil, err
	}

	return &request, nil
}

func (r *RequestRepo) GetRequestsByMessage(message string) ([]models.Requests, error) {
	rows, err := r.db.Db.Query(
		`SELECT request_id, request_message, user_id FROM requests WHERE request_message = $1`,
		message,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	requests := make([]models.Requests, 0)
	for rows.Next() {
		var request models.Requests
		err = rows.Scan(&request.Id, &request.Message, &request.UserId)
		if err != nil {
			return nil, err
		}

		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(requests) == 0 {
		return nil, err
	}

	return requests, nil
}

func (r *RequestRepo) GetRequestsByUserId(userId uint64) ([]models.Requests, error) {
	rows, err := r.db.Db.Query(
		`SELECT request_id, request_message, user_id FROM requests WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	requests := make([]models.Requests, 0)
	for rows.Next() {
		var request models.Requests
		err = rows.Scan(&request.Id, &request.Message, &request.UserId)
		if err != nil {
			return nil, err
		}

		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(requests) == 0 {
		return nil, err
	}

	return requests, nil
}
