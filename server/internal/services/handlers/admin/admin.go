package admins_service

import (
	"ural-hackaton/internal/config"
	admins_storage "ural-hackaton/internal/storage/repositories/admin"
)

type AdminService struct {
	repo *admins_storage.AdminRepo
	cfg  *config.Config
}

func Init(adminRepo *admins_storage.AdminRepo, cfg *config.Config) *AdminService {
	return &AdminService{
		repo: adminRepo,
		cfg:  cfg,
	}
}
