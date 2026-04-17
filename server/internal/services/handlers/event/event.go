package event_service

import (
	"time"
	"ural-hackaton/internal/config"
	event_dto "ural-hackaton/internal/dto/event"
	"ural-hackaton/internal/models"
	event_storage "ural-hackaton/internal/storage/repositories/event"
)

type EventService struct {
	repo *event_storage.EventRepo
	cfg  *config.Config
}

func Init(userRepo *event_storage.EventRepo, cfg *config.Config) *EventService {
	return &EventService{
		repo: userRepo,
		cfg:  cfg,
	}
}

func (s *EventService) CreateEvent(name string, description string, start time.Time, end time.Time, hub_id uint64, mentorId *uint64) error {

	eventDto := &event_dto.CreateEventDto{
		Name:        name,
		Description: description,
		StartTime:   start,
		EndTime:     end,
		HubId:       hub_id,
		MentorId:    mentorId,
	}

	err := s.repo.CreateEvent(eventDto)

	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetEventById(id uint64) (*models.Event, error) {
	event, err := s.repo.GetEventById(id)

	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) GetEventByName(name string) (*models.Event, error) {
	event, err := s.repo.GetEventByName(name)

	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) GetEventsByHubId(hubId uint64) ([]models.Event, error) {
	events, err := s.repo.GetEventsByHubId(hubId)

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	model, err := s.repo.GetAllEvents()

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *EventService) SearchEventsByName(query string) ([]models.Event, error) {
	model, err := s.repo.SearchEventsByName(query)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *EventService) UpdateEvent(eventName string, description string, start time.Time, end time.Time, hubId uint64, mentorId *uint64, id uint64) (*models.Event, error) {
	eventDto := &models.Event{
		Id:          id,
		EventName:   eventName,
		Description: description,
		StartTime:   start,
		EndTime:     end,
		HubId:       hubId,
		MentorId:    mentorId,
	}

	model, err := s.repo.UpdateEvent(eventDto)

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *EventService) DeleteEvent(id uint64) error {
	err := s.repo.DeleteEvent(id)

	if err != nil {
		return err
	}
	return nil
}
