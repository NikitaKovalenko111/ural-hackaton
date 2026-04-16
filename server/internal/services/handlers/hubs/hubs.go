package hub_service

import (
	"ural-hackaton/internal/config"
	hub_storage "ural-hackaton/internal/storage/repositories/hubs"
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
