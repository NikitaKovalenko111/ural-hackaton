package user_service

import (
	"strings"
	"ural-hackaton/internal/config"
	admin_dto "ural-hackaton/internal/dto/admin"
	mentor_dto "ural-hackaton/internal/dto/mentor"
	user_dto "ural-hackaton/internal/dto/user"
	"ural-hackaton/internal/models"
	admin_storage "ural-hackaton/internal/storage/repositories/admin"
	mentor_storage "ural-hackaton/internal/storage/repositories/mentor"
	user_storage "ural-hackaton/internal/storage/repositories/user"
)

type UserService struct {
	repo       *user_storage.UserRepo
	adminRepo  *admin_storage.AdminRepo
	mentorRepo *mentor_storage.MentorRepo
	cfg        *config.Config
}

func Init(userRepo *user_storage.UserRepo, adminRepo *admin_storage.AdminRepo, mentorRepo *mentor_storage.MentorRepo, cfg *config.Config) *UserService {
	return &UserService{
		repo:       userRepo,
		adminRepo:  adminRepo,
		mentorRepo: mentorRepo,
		cfg:        cfg,
	}
}

func (s *UserService) CreateUser(fullname string, role string, email string, telegram string, phone string, hubId *uint64) error {

	userDto := &user_dto.CreateUserDto{
		Fullname: fullname,
		Role:     role,
		Email:    email,
		Telegram: telegram,
		Phone:    phone,
		HubId:    hubId,
	}

	userID, err := s.repo.CreateUser(userDto)

	if err != nil {
		return err
	}

	roleNormalized := strings.ToLower(strings.TrimSpace(role))
	switch roleNormalized {
	case "admin":
		if s.adminRepo == nil {
			return nil
		}

		_, err = s.adminRepo.CreateAdmin(admin_dto.CreateAdminDto{UserId: userID})
		if err != nil {
			return err
		}
	case "mentor":
		if s.mentorRepo == nil {
			return nil
		}

		_, err = s.mentorRepo.CreateMentor(mentor_dto.CreateMentorDto{UserId: userID, HubId: hubId})
		if err != nil {
			return err
		}
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

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByTelegram(telegram string) (*models.User, error) {
	user, err := s.repo.GetUserByTelegram(telegram)

	if err != nil {
		return nil, err
	}
	return user, nil
}
