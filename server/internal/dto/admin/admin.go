package admin_dto

import "ural-hackaton/internal/models"

type AdminJoinUserDto struct {
	AdminId uint64 `json:"admin_id"`
	models.User
}

type CreateAdminDto struct {
	UserId uint64 `json:"user_id"`
}
