package hub_service

import (
	"ural-hackaton/internal/config"
	hub_dto "ural-hackaton/internal/dto/hub"
	"ural-hackaton/internal/models"
	hub_storage "ural-hackaton/internal/storage/repositories/hub"
)

type HubService struct {
	repo *hub_storage.HubRepo
	cfg  *config.Config
}

func Init(hubRepo *hub_storage.HubRepo, cfg *config.Config) *HubService {
	return &HubService{
		repo: hubRepo,
		cfg:  cfg,
	}
}

func (s *HubService) CreateHub(name string, address string) (*models.Hub, error) {
	hubDto := &hub_dto.CreateHubDto{
		Name:    name,
		Address: address,
	}

	model, err := s.repo.CreateHub(hubDto)

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *HubService) GetAllHubs() ([]models.Hub, error) {
	model, err := s.repo.GetAllHubs()

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *HubService) GetHubById(id uint64) (*models.Hub, error) {
	model, err := s.repo.GetHubById(id)

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *HubService) UpdateHub(hubName string, address string, status string, id uint64) (*models.Hub, error) {
	hubDto := &models.Hub{
		Id:      id,
		HubName: hubName,
		Address: address,
		Status:  status,
	}

	model, err := s.repo.UpdateHub(hubDto)

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *HubService) DeleteHub(id uint64) error {
	err := s.repo.DeleteHub(id)

	if err != nil {
		return err
	}
	return nil
}
