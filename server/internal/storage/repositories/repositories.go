package repositories

import (
	"ural-hackaton/internal/storage"
	user_storage "ural-hackaton/internal/storage/repositories/users"
)

type Repositories struct {
	UserRepository *user_storage.UserRepo
}

func InitRepositories(db *storage.Storage) *Repositories {
	userRepository := user_storage.Init(db)

	Repositories := Repositories{
		UserRepository: userRepository,
	}

	return &Repositories
}
