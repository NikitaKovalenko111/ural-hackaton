package users_storage

import (
	//"context"
	//"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
	user_storage_dto "ural-hackaton/internal/storage/repositories/users/dto"
)

type UserRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *user_storage_dto.CreateUserDto) {

}
