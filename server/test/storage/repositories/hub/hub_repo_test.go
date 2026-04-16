package hubs_storage_test

import (
	"database/sql/driver"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	hubsDto "ural-hackaton/internal/dto/hub"
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
	hubs_storage "ural-hackaton/internal/storage/repositories/hub"
)

func setupMockHub(t *testing.T) (*storage.Storage, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	st := storage.Init(db)
	return st, mock, func() { db.Close() }
}

func TestGetAllHubs_Success(t *testing.T) {
	st, mock, teardown := setupMockHub(t)
	defer teardown()

	rows := sqlmock.NewRows([]string{"hub_id", "hub_name"}).AddRow(1, "H1")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT hub_id, hub_name FROM hubs")).WillReturnRows(rows)

	repo := hubs_storage.Init(st)
	res, err := repo.GetAllHubs()

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, uint64(1), res[0].Id)
}

func TestGetAllHubs_QueryError(t *testing.T) {
	st, mock, teardown := setupMockHub(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT hub_id, hub_name FROM hubs")).WillReturnError(errors.New("db error"))

	repo := hubs_storage.Init(st)
	res, err := repo.GetAllHubs()

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetHubById_Success(t *testing.T) {
	st, mock, teardown := setupMockHub(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT hub_id, hub_name FROM hubs WHERE hub_id = $1")).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"hub_id", "hub_name"}).AddRow(2, "H2"))

	repo := hubs_storage.Init(st)
	res, err := repo.GetHubById(2)

	assert.NoError(t, err)
	assert.Equal(t, uint64(2), res.Id)
}

func TestCreateHub_Success(t *testing.T) {
	st, mock, teardown := setupMockHub(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO hubs (hub_name) VALUES ($1) RETURNING hub_id, hub_name")).WithArgs("NewHub").WillReturnRows(sqlmock.NewRows([]string{"hub_id", "hub_name"}).AddRow(10, "NewHub"))

	repo := hubs_storage.Init(st)
	created, err := repo.CreateHub(&hubsDto.CreateHubDto{Name: "NewHub"})

	assert.NoError(t, err)
	assert.Equal(t, uint64(10), created.Id)
	assert.Equal(t, "NewHub", created.HubName)
}

func TestUpdateHub_Success(t *testing.T) {
	st, mock, teardown := setupMockHub(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("UPDATE hubs SET hub_name = $1 WHERE hub_id = $2 RETURNING hub_id, hub_name")).WithArgs("Updated", 5).WillReturnRows(sqlmock.NewRows([]string{"hub_id", "hub_name"}).AddRow(5, "Updated"))

	repo := hubs_storage.Init(st)
	updated, err := repo.UpdateHub(&models.Hub{Id: 5, HubName: "Updated"})

	assert.NoError(t, err)
	assert.Equal(t, uint64(5), updated.Id)
}

func TestDeleteHub_Success(t *testing.T) {
	st, mock, teardown := setupMockHub(t)
	defer teardown()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM hubs WHERE hub_id = $1")).WithArgs(3).WillReturnResult(driver.ResultNoRows)

	repo := hubs_storage.Init(st)
	err := repo.DeleteHub(3)

	assert.NoError(t, err)
}
