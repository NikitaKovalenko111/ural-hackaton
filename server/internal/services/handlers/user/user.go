package users_service

import (
	"ural-hackaton/internal/config"
	user_dto "ural-hackaton/internal/dto/user"
	user_storage "ural-hackaton/internal/storage/repositories/user"
	"ural-hackaton/internal/types"
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

func (s *UserService) CreateUser(fullname string, role string) (*types.RequestStatus, error) {

	userDto := &user_dto.CreateUserDto{
		Fullname: fullname,
		Role:     role,
	}

	createdUser, err := s.repo.CreateUser(userDto)

	if err != nil {
		return nil, err
	}
}
