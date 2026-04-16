package admins_storage

import (
	admin_dto "ural-hackaton/internal/dto/admin"
	"ural-hackaton/internal/storage"
)

type AdminRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *AdminRepo {
	return &AdminRepo{db: db}
}

func (r *AdminRepo) GetAllAdmins() ([]admin_dto.AdminJoinUserDto, error) {
	rows, err := r.db.Db.Query(`
		SELECT admins.admin_id, users.user_id, users.fullname, users.user_role, users.email, users.telegram, users.phone
		FROM admins
		JOIN users ON admins.user_id = users.user_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []admin_dto.AdminJoinUserDto
	for rows.Next() {
		var admin admin_dto.AdminJoinUserDto
		err := rows.Scan(&admin.AdminId, &admin.Id, &admin.FullName, &admin.Role, &admin.Email, &admin.Telegram, &admin.Phone)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func (r *AdminRepo) GetAdminById(id uint64) (*admin_dto.AdminJoinUserDto, error) {
	var admin admin_dto.AdminJoinUserDto

	err := r.db.Db.QueryRow(
<<<<<<< HEAD
		`SELECT admin_id FROM admins WHERE admin_id = $1`,
		id,
	).Scan(&admin.AdminId)
=======
		`SELECT admins.admin_id, users.user_id, users.fullname, users.user_role, users.email, users.telegram, users.phone
		 FROM admins
		 JOIN users ON admins.user_id = users.user_id
		 WHERE admins.admin_id = $1`,
		id,
	).Scan(&admin.AdminId, &admin.Id, &admin.FullName, &admin.Role, &admin.Email, &admin.Telegram, &admin.Phone)
>>>>>>> 99e69862683c0673a6a86170ad35b8c424f1a8d0

	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *AdminRepo) CreateAdmin(admin admin_dto.CreateAdminDto) (*admin_dto.AdminJoinUserDto, error) {
	var createdAdmin admin_dto.AdminJoinUserDto

	err := r.db.Db.QueryRow(
		`INSERT INTO admins (user_id) VALUES ($1) RETURNING admin_id, user_id`,
		admin.UserId,
	).Scan(&createdAdmin.AdminId, &createdAdmin.Id)

	if err != nil {
		return nil, err
	}

	return &createdAdmin, nil
}

func (r *AdminRepo) DeleteAdmin(id uint64) error {
	_, err := r.db.Db.Exec(
		`DELETE FROM admins WHERE admin_id = $1`,
		id,
	)

	return err
}

func (r *AdminRepo) GetAdminByFullname(fullname string) (*admin_dto.AdminJoinUserDto, error) {
	var admin admin_dto.AdminJoinUserDto

	err := r.db.Db.QueryRow(
		`SELECT admins.admin_id, users.user_id, users.fullname, users.user_role, users.email, users.telegram, users.phone
		 FROM admins
		 JOIN users ON admins.user_id = users.user_id
		 WHERE users.fullname = $1`,
		fullname,
	).Scan(&admin.AdminId, &admin.Id, &admin.FullName, &admin.Role, &admin.Email, &admin.Telegram, &admin.Phone)

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
