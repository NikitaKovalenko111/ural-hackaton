package requests_storage_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"ural-hackaton/internal/storage"
	requests_storage "ural-hackaton/internal/storage/repositories/request"
)

func setupMockRequest(t *testing.T) (*storage.Storage, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	st := storage.Init(db)
	return st, mock, func() { db.Close() }
}

func TestCreateRequest_Success(t *testing.T) {
	st, mock, teardown := setupMockRequest(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO requests (message, user_id) VALUES ($1, $2) RETURNING request_id")).WithArgs("hi", 1).WillReturnRows(sqlmock.NewRows([]string{"request_id"}).AddRow(5))

	repo := requests_storage.Init(st)
	err := repo.CreateRequest("hi", 1)

	assert.NoError(t, err)
}

func TestGetRequestById_Success(t *testing.T) {
	st, mock, teardown := setupMockRequest(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT request_id, message, user_id FROM requests WHERE request_id = $1")).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"request_id", "message", "user_id"}).AddRow(2, "hi", 1))

	repo := requests_storage.Init(st)
	res, err := repo.GetRequestById(2)

	assert.NoError(t, err)
	assert.Equal(t, uint64(2), res.Id)
}

func TestGetRequestsByMessage_Success(t *testing.T) {
	st, mock, teardown := setupMockRequest(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT request_id, message, user_id FROM requests WHERE message = $1")).WithArgs("hi").WillReturnRows(sqlmock.NewRows([]string{"request_id", "message", "user_id"}).AddRow(1, "hi", 1))

	repo := requests_storage.Init(st)
	res, err := repo.GetRequestsByMessage("hi")

	assert.NoError(t, err)
	assert.Len(t, res, 1)
}

func TestGetRequestsByUserId_Success(t *testing.T) {
	st, mock, teardown := setupMockRequest(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT request_id, message, user_id FROM requests WHERE user_id = $1")).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"request_id", "message", "user_id"}).AddRow(1, "hi", 1))

	repo := requests_storage.Init(st)
	res, err := repo.GetRequestsByUserId(1)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
}
