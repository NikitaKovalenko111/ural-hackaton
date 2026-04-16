package admins_storage_test

import (
	"database/sql/driver"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	admin_dto "ural-hackaton/internal/dto/admin"
	"ural-hackaton/internal/storage"
	admins_storage "ural-hackaton/internal/storage/repositories/admin"
)

func setupMock(t *testing.T) (*storage.Storage, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	st := storage.Init(db)

	teardown := func() { db.Close() }
	return st, mock, teardown
}

func TestGetAllAdmins_Success(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	rows := sqlmock.NewRows([]string{"admin_id", "fullname"}).AddRow(1, "Ivan")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT admin_id, fullname FROM admins JOIN users ON admins.user_id = users.user_id")).WillReturnRows(rows)

	repo := admins_storage.Init(st)
	res, err := repo.GetAllAdmins()

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, uint64(1), res[0].AdminId)
}

func TestGetAllAdmins_QueryError(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT admin_id, fullname FROM admins JOIN users ON admins.user_id = users.user_id")).WillReturnError(errors.New("db error"))

	repo := admins_storage.Init(st)
	res, err := repo.GetAllAdmins()

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetAdminById_Success(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT admin_id, fullname FROM admins WHERE admin_id = $1")).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"admin_id", "fullname"}).AddRow(2, "Ivan"))

	repo := admins_storage.Init(st)
	res, err := repo.GetAdminById(2)

	assert.NoError(t, err)
	assert.Equal(t, uint64(2), res.AdminId)
}

func TestGetAdminById_NotFound(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT admin_id, fullname FROM admins WHERE admin_id = $1")).WithArgs(5).WillReturnError(errors.New("no rows"))

	repo := admins_storage.Init(st)
	res, err := repo.GetAdminById(5)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestCreateAdmin_Success(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO admins (user_id) VALUES ($1) RETURNING admin_id, user_id")).WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"admin_id", "user_id"}).AddRow(10, 3))

	repo := admins_storage.Init(st)
	created, err := repo.CreateAdmin(admin_dto.CreateAdminDto{UserId: 3})

	assert.NoError(t, err)
	assert.Equal(t, uint64(10), created.AdminId)
	assert.Equal(t, uint64(3), created.Id)
}

func TestCreateAdmin_InsertError(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO admins (user_id) VALUES ($1) RETURNING admin_id, user_id")).WithArgs(3).WillReturnError(errors.New("insert failed"))

	repo := admins_storage.Init(st)
	created, err := repo.CreateAdmin(admin_dto.CreateAdminDto{UserId: 3})

	assert.Error(t, err)
	assert.Nil(t, created)
}

func TestDeleteAdmin_Success(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM admins WHERE admin_id = $1")).WithArgs(4).WillReturnResult(driver.ResultNoRows)

	repo := admins_storage.Init(st)
	err := repo.DeleteAdmin(4)

	assert.NoError(t, err)
}

func TestDeleteAdmin_ExecError(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM admins WHERE admin_id = $1")).WithArgs(4).WillReturnError(errors.New("delete failed"))

	repo := admins_storage.Init(st)
	err := repo.DeleteAdmin(4)

	assert.Error(t, err)
}

func TestGetAdminByFullname_Success(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT admin_id, fullname FROM admins WHERE fullname = $1")).WithArgs("Ivan Ivanov").WillReturnRows(sqlmock.NewRows([]string{"admin_id", "fullname"}).AddRow(7, "Ivan Ivanov"))

	repo := admins_storage.Init(st)
	res, err := repo.GetAdminByFullname("Ivan Ivanov")

	assert.NoError(t, err)
	assert.Equal(t, uint64(7), res.AdminId)
}

func TestGetAdminByFullname_NoRows(t *testing.T) {
	st, mock, teardown := setupMock(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT admin_id, fullname FROM admins WHERE fullname = $1")).WithArgs("Nobody").WillReturnError(errors.New("not found"))

	repo := admins_storage.Init(st)
	res, err := repo.GetAdminByFullname("Nobody")

	assert.Error(t, err)
	assert.Nil(t, res)
}
