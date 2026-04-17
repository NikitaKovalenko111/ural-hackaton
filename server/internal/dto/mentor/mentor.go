package mentor_dto

import "ural-hackaton/internal/models"

type CreateMentorDto struct {
	UserId uint64  `json:"user_id"`
	HubId  *uint64 `json:"hub_id"`
}

type MentorJoinUserDto struct {
	MentorId uint64  `json:"mentor_id"`
	HubId    *uint64 `json:"hub_id,omitempty"`
	models.User
}
