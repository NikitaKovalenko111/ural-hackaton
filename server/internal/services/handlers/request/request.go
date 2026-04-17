package request_service

import (
	"ural-hackaton/internal/config"
	"ural-hackaton/internal/models"
	request_storage "ural-hackaton/internal/storage/repositories/request"
)

type RequestService struct {
	repo *request_storage.RequestRepo
	cfg  *config.Config
}

func Init(requestRepo *request_storage.RequestRepo, cfg *config.Config) *RequestService {
	return &RequestService{
		repo: requestRepo,
		cfg:  cfg,
	}
}

func (s *RequestService) CreateRequest(message string, userId uint64, mentorId *uint64) error {

	err := s.repo.CreateRequest(message, userId, mentorId)

	if err != nil {
		return err
	}
	return nil
}

func (s *RequestService) GetRequestById(id uint64) (*models.Requests, error) {

	request, err := s.repo.GetRequestById(id)

	if err != nil {
		return nil, err
	}
	return request, nil
}

func (s *RequestService) GetRequestsByMessage(message string) ([]models.Requests, error) {

	request, err := s.repo.GetRequestsByMessage(message)

	if err != nil {
		return nil, err
	}
	return request, nil
}

func (s *RequestService) GetRequestsByUserId(userId uint64) ([]models.Requests, error) {

	request, err := s.repo.GetRequestsByUserId(userId)

	if err != nil {
		return nil, err
	}
	return request, nil
}

func (s *RequestService) GetRequestsByMentorId(mentorId uint64) ([]models.Requests, error) {

	request, err := s.repo.GetRequestsByMentorId(mentorId)

	if err != nil {
		return nil, err
	}
	return request, nil
}
