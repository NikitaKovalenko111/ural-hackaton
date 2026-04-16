package repositories

import (
	"ural-hackaton/internal/storage"
	admin_storage "ural-hackaton/internal/storage/repositories/admin"
	hub_storage "ural-hackaton/internal/storage/repositories/hub"
	mentor_storage "ural-hackaton/internal/storage/repositories/mentor"
	requests_storage "ural-hackaton/internal/storage/repositories/request"
	user_storage "ural-hackaton/internal/storage/repositories/user"
)

type Repositories struct {
	UserRepository    *user_storage.UserRepo
	HubRepository     *hub_storage.HubRepo
	AdminRepository   *admin_storage.AdminRepo
	RequestRepository *requests_storage.RequestRepo
	MentorRepository  *mentor_storage.MentorRepo
}

func InitRepositories(db *storage.Storage) *Repositories {
	return &Repositories{
		UserRepository:    user_storage.Init(db),
		HubRepository:     hub_storage.Init(db),
		AdminRepository:   admin_storage.Init(db),
		RequestRepository: requests_storage.Init(db),
		MentorRepository:  mentor_storage.Init(db),
	}
}
