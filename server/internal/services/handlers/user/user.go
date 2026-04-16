package users_service

import (
	"ural-hackaton/internal/config"
	user_dto "ural-hackaton/internal/dto/user"
	"ural-hackaton/internal/models"
	user_storage "ural-hackaton/internal/storage/repositories/user"
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

	userDto := &user_dto.CreateUserDto{
		Fullname: fullname,
		Role:     role,
	}

	status, err := s.repo.CreateUser(userDto)

	if err != nil {
		return nil, err
	}
	return status, nil
}

func (s *UserService) GetUserById(id uint64) (*models.User, *fiber.Error) {
	user, err := s.repo.GetUserById(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByFullname(fullname string) (*models.User, *fiber.Error) {
	user, err := s.repo.GetUserByFullname(fullname)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsersByRole(role string) ([]models.User, *fiber.Error) {
	user, err := s.repo.GetUsersByRole(role)

	if err != nil {
		return nil, err
	}
	return user, nil
}
