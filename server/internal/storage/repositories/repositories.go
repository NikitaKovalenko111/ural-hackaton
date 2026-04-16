package repositories

import "ural-hackaton/internal/storage"

type Repositories struct {
}

func InitRepositories(db *storage.Storage) *Repositories {
	Repositories := Repositories{}

	return &Repositories
}
