package mentor_service

import (
	"ural-hackaton/internal/config"
	mentor_dto "ural-hackaton/internal/dto/mentor"
	mentor_storage "ural-hackaton/internal/storage/repositories/mentor"
)

type MentorService struct {
	repo *mentor_storage.MentorRepo
	cfg  *config.Config
}

func Init(hubRepo *mentor_storage.MentorRepo, cfg *config.Config) *MentorService {
	return &MentorService{
		repo: hubRepo,
		cfg:  cfg,
	}
}

func (s *MentorService) CreateMentor(userId uint64, hubId *uint64) (*mentor_dto.MentorJoinUserDto, error) {
	mentorDto := mentor_dto.CreateMentorDto{
		UserId: userId,
		HubId:  hubId,
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

func (s *MentorService) GetMentorByUserId(userId uint64) (*mentor_dto.MentorJoinUserDto, error) {
	mentor, err := s.repo.GetMentorByUserId(userId)
	if err != nil {
		return nil, err
	}

	return mentor, nil
}

func (s *MentorService) GetAllMentors() ([]mentor_dto.MentorJoinUserDto, error) {
	mentors, err := s.repo.GetAllMentors()
	if err != nil {
		return nil, err
	}

	return mentors, nil
}
