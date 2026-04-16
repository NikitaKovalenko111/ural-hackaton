package hubs_service

import (
	"ural-hackaton/internal/config"
	hubsStorageDto "ural-hackaton/internal/dto/hub"
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

func (s *HubService) CreateHub(name string) (*models.Hub, error) {
	hubDto := &hubsStorageDto.CreateHubDto{
		Name: name,
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

func (s *HubService) UpdateHub(hubName string, id uint64) (*models.Hub, error) {
	hubDto := &models.Hub{
		Id:      id,
		HubName: hubName,
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
