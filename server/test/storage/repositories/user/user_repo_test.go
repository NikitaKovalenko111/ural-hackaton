package users_storage_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	usersDto "ural-hackaton/internal/dto/user"
	"ural-hackaton/internal/storage"
	users_storage "ural-hackaton/internal/storage/repositories/user"
)

func setupMockUser(t *testing.T) (*storage.Storage, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	st := storage.Init(db)
	return st, mock, func() { db.Close() }
}

func TestCreateUser_Success(t *testing.T) {
	st, mock, teardown := setupMockUser(t)
	defer teardown()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (fullname, user_role) VALUES ($1, $2)")).WithArgs("Ivan", "admin").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := users_storage.Init(st)
	err := repo.CreateUser(&usersDto.CreateUserDto{Fullname: "Ivan", Role: "admin"})

	assert.NoError(t, err)
}

func TestGetUserById_Success(t *testing.T) {
	st, mock, teardown := setupMockUser(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT user_id, fullname, user_role FROM users WHERE user_id = $1")).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"user_id", "fullname", "user_role"}).AddRow(2, "Ivan", "admin"))

	repo := users_storage.Init(st)
	res, err := repo.GetUserById(2)

	assert.NoError(t, err)
	assert.Equal(t, uint64(2), res.Id)
}

func TestGetUserByFullname_Success(t *testing.T) {
	st, mock, teardown := setupMockUser(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT user_id, fullname, user_role FROM users WHERE fullname = $1")).WithArgs("Ivan").WillReturnRows(sqlmock.NewRows([]string{"user_id", "fullname", "user_role"}).AddRow(3, "Ivan", "admin"))

	repo := users_storage.Init(st)
	res, err := repo.GetUserByFullname("Ivan")

	assert.NoError(t, err)
	assert.Equal(t, uint64(3), res.Id)
}

func TestGetUsersByRole_Success(t *testing.T) {
	st, mock, teardown := setupMockUser(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT user_id, fullname, user_role FROM users WHERE user_role = $1")).WithArgs("admin").WillReturnRows(sqlmock.NewRows([]string{"user_id", "fullname", "user_role"}).AddRow(1, "Ivan", "admin"))

	repo := users_storage.Init(st)
	res, err := repo.GetUsersByRole("admin")

	assert.NoError(t, err)
	assert.Len(t, res, 1)
}
