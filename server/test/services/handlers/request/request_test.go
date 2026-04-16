package requests_service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"ural-hackaton/internal/models"
	request_service "ural-hackaton/internal/services/handlers/request"
)

// manual mock for request repo
type mockRequestRepo struct {
	createFn       func(string, uint64) error
	getByIdFn      func(uint64) (*models.Requests, error)
	getByMessageFn func(string) ([]models.Requests, error)
	getByUserIdFn  func(uint64) ([]models.Requests, error)
}

func (m *mockRequestRepo) CreateRequest(message string, userId uint64) error {
	return m.createFn(message, userId)
}
func (m *mockRequestRepo) GetRequestById(id uint64) (*models.Requests, error) { return m.getByIdFn(id) }
func (m *mockRequestRepo) GetRequestsByMessage(message string) ([]models.Requests, error) {
	return m.getByMessageFn(message)
}
func (m *mockRequestRepo) GetRequestsByUserId(userId uint64) ([]models.Requests, error) {
	return m.getByUserIdFn(userId)
}

func TestCreateRequest_Success(t *testing.T) {
	mock := &mockRequestRepo{}
	mock.createFn = func(message string, userId uint64) error { return nil }

	svc := request_service.Init(mock, nil)
	err := svc.CreateRequest("hi", 1)

	assert.NoError(t, err)
}

func TestCreateRequest_Error(t *testing.T) {
	mock := &mockRequestRepo{}
	mock.createFn = func(message string, userId uint64) error { return errors.New("db") }

	svc := request_service.Init(mock, nil)
	err := svc.CreateRequest("hi", 1)

	assert.Error(t, err)
}
