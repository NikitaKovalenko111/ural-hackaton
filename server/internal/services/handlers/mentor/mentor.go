package mentor_service

import (
	"ural-hackaton/internal/config"
	mentor_dto "ural-hackaton/internal/dto/mentor"
	mentor_storage "ural-hackaton/internal/storage/repositories/mentor"
)

type MentorService struct {
	repo mentor_storage.MentorRepoI
	cfg  *config.Config
}

func Init(hubRepo mentor_storage.MentorRepoI, cfg *config.Config) *MentorService {
	return &MentorService{
		repo: hubRepo,
		cfg:  cfg,
	}
}

func (s *MentorService) CreateMentor(userId uint64) (*mentor_dto.MentorJoinUserDto, error) {
	mentorDto := mentor_dto.CreateMentorDto{
		UserId: userId,
	}

	mentor, err := s.repo.CreateMentor(mentorDto)

	if err != nil {
		return nil, err
	}
	return mentor, nil
}

func (s *MentorService) GetMentorById(id uint64) (*mentor_dto.MentorJoinUserDto, error) {
	mentor, err := s.repo.GetMentorById(id)

	if err != nil {
		return nil, err
	}
	return mentor, nil
}

func (s *MentorService) GetMentorByFullname(fullname string) (*mentor_dto.MentorJoinUserDto, error) {
	mentor, err := s.repo.GetMentorByFullname(fullname)

	if err != nil {
		return nil, err
	}
	return mentor, nil
}
