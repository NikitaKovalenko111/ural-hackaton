package repositories

import (
	"ural-hackaton/internal/storage"
	admin_storage "ural-hackaton/internal/storage/repositories/admin"
	hub_storage "ural-hackaton/internal/storage/repositories/hub"
	user_storage "ural-hackaton/internal/storage/repositories/user"
)

type Repositories struct {
	UserRepository  *user_storage.UserRepo
	HubRepository   *hub_storage.HubRepo
	AdminRepository *admin_storage.AdminsRepo
}

func InitRepositories(db *storage.Storage) *Repositories {
	return &Repositories{
		UserRepository:  user_storage.Init(db),
		HubRepository:   hub_storage.Init(db),
		AdminRepository: admin_storage.Init(db),
	}
}
