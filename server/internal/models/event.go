package models

import "time"

type Event struct {
	Id          uint64    `json:"event_id"`
	EventName   string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start"`
	EndTime     time.Time `json:"end"`
	HubId       uint64    `json:"hub_id"`
}
