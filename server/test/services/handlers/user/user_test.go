package users_service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	usersDto "ural-hackaton/internal/dto/user"
	"ural-hackaton/internal/models"
	user_service "ural-hackaton/internal/services/handlers/user"
)

// manual mock for user repo
type mockUserRepo struct {
	createFn        func(*usersDto.CreateUserDto) error
	getByIdFn       func(uint64) (*models.User, error)
	getByFullnameFn func(string) (*models.User, error)
	getByRoleFn     func(string) ([]models.User, error)
}

func (m *mockUserRepo) CreateUser(u *usersDto.CreateUserDto) error  { return m.createFn(u) }
func (m *mockUserRepo) GetUserById(id uint64) (*models.User, error) { return m.getByIdFn(id) }
func (m *mockUserRepo) GetUserByFullname(name string) (*models.User, error) {
	return m.getByFullnameFn(name)
}
func (m *mockUserRepo) GetUsersByRole(role string) ([]models.User, error) { return m.getByRoleFn(role) }

func TestCreateUser_Success(t *testing.T) {
	mock := &mockUserRepo{}
	mock.createFn = func(u *usersDto.CreateUserDto) error { return nil }

	svc := user_service.Init(mock, nil)
	err := svc.CreateUser("Ivan", "admin")

	assert.NoError(t, err)
}

func TestCreateUser_Error(t *testing.T) {
	mock := &mockUserRepo{}
	mock.createFn = func(u *usersDto.CreateUserDto) error { return errors.New("db") }

	svc := user_service.Init(mock, nil)
	err := svc.CreateUser("Ivan", "admin")

	assert.Error(t, err)
}
