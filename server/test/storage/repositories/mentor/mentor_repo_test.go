package mentor_storage_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	mentor_dto "ural-hackaton/internal/dto/mentor"
	"ural-hackaton/internal/storage"
	mentor_storage "ural-hackaton/internal/storage/repositories/mentor"
)

func setupMockMentor(t *testing.T) (*storage.Storage, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	st := storage.Init(db)
	return st, mock, func() { db.Close() }
}

func TestCreateMentor_Success(t *testing.T) {
	st, mock, teardown := setupMockMentor(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO mentors (user_id) VALUES ($1) RETURNING mentor_id, user_id")).WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"mentor_id", "user_id"}).AddRow(11, 3))

	repo := mentor_storage.Init(st)
	res, err := repo.CreateMentor(mentor_dto.CreateMentorDto{UserId: 3})

	assert.NoError(t, err)
	assert.Equal(t, uint64(11), res.MentorId)
	assert.Equal(t, uint64(3), res.Id)
}

func TestCreateMentor_Error(t *testing.T) {
	st, mock, teardown := setupMockMentor(t)
	defer teardown()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO mentors (user_id) VALUES ($1) RETURNING mentor_id, user_id")).WithArgs(3).WillReturnError(errors.New("insert failed"))

	repo := mentor_storage.Init(st)
	res, err := repo.CreateMentor(mentor_dto.CreateMentorDto{UserId: 3})

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetMentorById_Success(t *testing.T) {
	st, mock, teardown := setupMockMentor(t)
	defer teardown()

	query := `SELECT mentors.mentor_id, users.user_id, users.user_fullname, users.user_role
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE mentors.mentor_id = $1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"mentor_id", "user_id", "user_fullname", "user_role"}).AddRow(2, 2, "M", "mentor"))

	repo := mentor_storage.Init(st)
	res, err := repo.GetMentorById(2)

	assert.NoError(t, err)
	assert.Equal(t, uint64(2), res.MentorId)
}

func TestGetMentorByFullname_Success(t *testing.T) {
	st, mock, teardown := setupMockMentor(t)
	defer teardown()

	query := `SELECT mentors.mentor_id, users.user_id, users.user_fullname, users.user_role
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE users.user_fullname = $1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("Ivan").WillReturnRows(sqlmock.NewRows([]string{"mentor_id", "user_id", "user_fullname", "user_role"}).AddRow(7, 7, "Ivan", "mentor"))

	repo := mentor_storage.Init(st)
	res, err := repo.GetMentorByFullname("Ivan")

	assert.NoError(t, err)
	assert.Equal(t, uint64(7), res.MentorId)
}
