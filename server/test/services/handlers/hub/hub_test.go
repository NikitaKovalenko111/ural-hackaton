package hubs_service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	hubsDto "ural-hackaton/internal/dto/hub"
	"ural-hackaton/internal/models"
	hub_service "ural-hackaton/internal/services/handlers/hub"
)

// manual mock for hub repo
type mockHubRepo struct {
	createFn  func(*hubsDto.CreateHubDto) (*models.Hub, error)
	getAllFn  func() ([]models.Hub, error)
	getByIdFn func(uint64) (*models.Hub, error)
	updateFn  func(*models.Hub) (*models.Hub, error)
	deleteFn  func(uint64) error
}

func (m *mockHubRepo) CreateHub(h *hubsDto.CreateHubDto) (*models.Hub, error) { return m.createFn(h) }
func (m *mockHubRepo) GetAllHubs() ([]models.Hub, error)                      { return m.getAllFn() }
func (m *mockHubRepo) GetHubById(id uint64) (*models.Hub, error)              { return m.getByIdFn(id) }
func (m *mockHubRepo) UpdateHub(h *models.Hub) (*models.Hub, error)           { return m.updateFn(h) }
func (m *mockHubRepo) DeleteHub(id uint64) error                              { return m.deleteFn(id) }

func TestCreateHub_Success(t *testing.T) {
	mock := &mockHubRepo{}
	mock.createFn = func(h *hubsDto.CreateHubDto) (*models.Hub, error) {
		return &models.Hub{Id: 1, HubName: h.Name}, nil
	}

	svc := hub_service.Init(mock, nil)
	res, err := svc.CreateHub("MyHub")

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, uint64(1), res.Id)
}

func TestCreateHub_Error(t *testing.T) {
	mock := &mockHubRepo{}
	mock.createFn = func(h *hubsDto.CreateHubDto) (*models.Hub, error) { return nil, errors.New("db") }

	svc := hub_service.Init(mock, nil)
	res, err := svc.CreateHub("MyHub")

	assert.Error(t, err)
	assert.Nil(t, res)
}
