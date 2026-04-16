package admin_service

import (
	"ural-hackaton/internal/config"
	admin_storage "ural-hackaton/internal/storage/repositories/admins"
)

type AdminService struct {
	repo *admin_storage.AdminRepo
	cfg  *config.Config
}

func Init(adminRepo *admin_storage.UserRepo, cfg *config.Config) *AdminService {
	return &AdminService{
		repo: adminRepo,
		cfg:  cfg,
	}
}
