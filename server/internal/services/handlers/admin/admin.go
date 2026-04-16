package admins_service

import (
	"ural-hackaton/internal/config"
	admin_dto "ural-hackaton/internal/dto/admin"
	admin_storage "ural-hackaton/internal/storage/repositories/admin"
)

type AdminService struct {
	repo admin_storage.AdminRepoI
	cfg  *config.Config
}

func Init(adminRepo admin_storage.AdminRepoI, cfg *config.Config) *AdminService {
	return &AdminService{
		repo: adminRepo,
		cfg:  cfg,
	}
}

func (s *AdminService) CreateAdmin(userId uint64) (*admin_dto.AdminJoinUserDto, error) {
	adminDto := admin_dto.CreateAdminDto{
		UserId: userId,
	}

	admin, err := s.repo.CreateAdmin(adminDto)

	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *AdminService) GetAllAdmins() ([]admin_dto.AdminJoinUserDto, error) {
	admin, err := s.repo.GetAllAdmins()

	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *AdminService) GetAdminById(id uint64) (*admin_dto.AdminJoinUserDto, error) {
	admin, err := s.repo.GetAdminById(id)

	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *AdminService) DeleteAdmin(id uint64) error {
	err := s.repo.DeleteAdmin(id)

	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) GetAdminByFullname(fullname string) (*admin_dto.AdminJoinUserDto, error) {
	admin, err := s.repo.GetAdminByFullname(fullname)

	if err != nil {
		return nil, err
	}
	return admin, nil
}
