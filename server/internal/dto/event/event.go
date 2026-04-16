package event_dto

import "time"

type CreateEventDto struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start"`
	EndTime     time.Time `json:"end"`
	HubId       uint64    `json:"hub_id"`
}
