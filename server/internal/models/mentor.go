package models

type Mentor struct {
	Id     uint64  `json:"mentor_id"`
	UserId uint64  `json:"user_id"`
	HubId  *uint64 `json:"hub_id,omitempty"`
}
