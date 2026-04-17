package event_dto

import "time"

type CreateEventDto struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	HubId       uint64    `json:"hub_id"`
	MentorId    *uint64   `json:"mentor_id"`
}
