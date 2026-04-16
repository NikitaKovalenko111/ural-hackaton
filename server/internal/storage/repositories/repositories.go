package repositories

import (
	"ural-hackaton/internal/storage"
	admin_storage "ural-hackaton/internal/storage/repositories/admins"
	hub_storage "ural-hackaton/internal/storage/repositories/hubs"
	user_storage "ural-hackaton/internal/storage/repositories/users"
)

type Repositories struct {
	UserRepository  *user_storage.UserRepo
	HubRepository   *hub_storage.HubRepo
	AdminRepository *admin_storage.AdminRepo
}

func InitRepositories(db *storage.Storage) *Repositories {
	return &Repositories{
		UserRepository:  user_storage.Init(db),
		HubRepository:   hub_storage.Init(db),
		AdminRepository: admin_storage.Init(db),
	}
}
