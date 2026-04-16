package user_service

import (
	"ural-hackaton/internal/config"
	user_dto "ural-hackaton/internal/dto/user"
	"ural-hackaton/internal/models"
	user_storage "ural-hackaton/internal/storage/repositories/user"
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

func (s *UserService) CreateUser(fullname string, role string, email string, telegram string) error {

	userDto := &user_dto.CreateUserDto{
		Fullname: fullname,
		Role:     role,
		Email:    email,
		Telegram: telegram,
	}

	err := s.repo.CreateUser(userDto)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserById(id uint64) (*models.User, error) {
	user, err := s.repo.GetUserById(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByFullname(fullname string) (*models.User, error) {
	user, err := s.repo.GetUserByFullname(fullname)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsersByRole(role string) ([]models.User, error) {
	user, err := s.repo.GetUsersByRole(role)

	if err != nil {
		return nil, err
	}
	return user, nil
}
