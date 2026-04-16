package user_service

import (
	"ural-hackaton/internal/config"
	user_dto "ural-hackaton/internal/dto"
	user_storage "ural-hackaton/internal/storage/repositories/users"
	"ural-hackaton/internal/types"

	"github.com/gofiber/fiber/v2"
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

func (s *UserService) CreateUser(fullname string, role string) (*types.RequestStatus, *fiber.Error) {

	userDto := &user_dto.CreateUserDto{}

	createdUser, err := s.repo.CreateUser()
}
