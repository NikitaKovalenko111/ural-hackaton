package requests_storage

import (
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
)

type RequestRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *RequestRepo {
	return &RequestRepo{db: db}
}

// RequestRepoI defines repository interface for requests
type RequestRepoI interface {
	CreateRequest(message string, userId uint64) error
	GetRequestById(id uint64) (*models.Requests, error)
	GetRequestsByMessage(message string) ([]models.Requests, error)
	GetRequestsByUserId(userId uint64) ([]models.Requests, error)
}

func (r *RequestRepo) CreateRequest(message string, userId uint64) error {

	err := r.db.Db.QueryRow(
		`INSERT INTO requests (message, user_id) VALUES ($1, $2) RETURNING request_id`,
		message, userId,
	).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *RequestRepo) GetRequestById(id uint64) (*models.Requests, error) {
	var request models.Requests

	err := r.db.Db.QueryRow(
		`SELECT request_id, message, user_id FROM requests WHERE request_id = $1`,
		id,
	).Scan(&request.Id, &request.Message, &request.UserId)

	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (r *RequestRepo) GetRequestsByMessage(message string) ([]models.Requests, error) {
	rows, err := r.db.Db.Query(
		`SELECT request_id, message, user_id FROM requests WHERE message = $1`,
		message,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []models.Requests
	for rows.Next() {
		var req models.Requests
		err = rows.Scan(&req.Id, &req.Message, &req.UserId)
		if err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}

	return requests, nil
}

func (r *RequestRepo) GetRequestsByUserId(userId uint64) ([]models.Requests, error) {
	rows, err := r.db.Db.Query(
		`SELECT request_id, message, user_id FROM requests WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []models.Requests
	for rows.Next() {
		var req models.Requests
		err = rows.Scan(&req.Id, &req.Message, &req.UserId)
		if err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}

	return requests, nil
}
