package user_service

import (
	"ural-hackaton/internal/config"
	user_storage "ural-hackaton/internal/storage/repositories/users"
)

type UserService struct {
	repo *user_storage.UserRepo
	cfg  *config.Config
}

func Init(userRepo *user_storage.UserRepo, cfg *config.Config) *UserService {
	return &UserService{
		repo: userRepo,
		cfg:  cfg,
	}
}
