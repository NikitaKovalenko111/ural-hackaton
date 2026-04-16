package mentor_dto

import "ural-hackaton/internal/models"

type CreateMentorDto struct {
	UserId uint64 `json:"user_id"`
}

type MentorJoinUserDto struct {
	MentorId uint64 `json:"mentor_id"`
	models.User
}
