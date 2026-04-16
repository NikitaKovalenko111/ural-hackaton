package requests_service

import (
	"ural-hackaton/internal/config"
	request_storage "ural-hackaton/internal/storage/repositories/request"
)

type ReauestService struct {
	repo *request_storage.RequestRepo
	cfg  *config.Config
}

func Init(requestRepo *request_storage.RequestRepo, cfg *config.Config) *ReauestService {
	return &ReauestService{
		repo: requestRepo,
		cfg:  cfg,
	}
}
