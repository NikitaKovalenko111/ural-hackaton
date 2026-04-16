package mentor_service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	mentor_dto "ural-hackaton/internal/dto/mentor"
	mentor_service "ural-hackaton/internal/services/handlers/mentor"
)

// manual mock for mentor repo
type mockMentorRepo struct {
	createFn        func(mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error)
	getByIdFn       func(uint64) (*mentor_dto.MentorJoinUserDto, error)
	getByFullnameFn func(string) (*mentor_dto.MentorJoinUserDto, error)
}

func (m *mockMentorRepo) CreateMentor(md mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error) {
	return m.createFn(md)
}
func (m *mockMentorRepo) GetMentorById(id uint64) (*mentor_dto.MentorJoinUserDto, error) {
	return m.getByIdFn(id)
}
func (m *mockMentorRepo) GetMentorByFullname(name string) (*mentor_dto.MentorJoinUserDto, error) {
	return m.getByFullnameFn(name)
}

func TestCreateMentor_Success(t *testing.T) {
	mock := &mockMentorRepo{}
	mock.createFn = func(md mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error) {
		return &mentor_dto.MentorJoinUserDto{MentorId: 1}, nil
	}

	svc := mentor_service.Init(mock, nil)
	res, err := svc.CreateMentor(1)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, uint64(1), res.MentorId)
}

func TestCreateMentor_Error(t *testing.T) {
	mock := &mockMentorRepo{}
	mock.createFn = func(md mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error) {
		return nil, errors.New("db")
	}

	svc := mentor_service.Init(mock, nil)
	res, err := svc.CreateMentor(1)

	assert.Error(t, err)
	assert.Nil(t, res)
}
