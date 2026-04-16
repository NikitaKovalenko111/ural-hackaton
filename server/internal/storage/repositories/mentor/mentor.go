package mentor_storage

import (
	mentor_dto "ural-hackaton/internal/dto/mentor"
	"ural-hackaton/internal/storage"
)

type MentorRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *MentorRepo {
	return &MentorRepo{
		db: db,
	}
}

// MentorRepoI defines interface for mentor repository
type MentorRepoI interface {
	CreateMentor(mentor mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error)
	GetMentorById(id uint64) (*mentor_dto.MentorJoinUserDto, error)
	GetMentorByFullname(fullname string) (*mentor_dto.MentorJoinUserDto, error)
}

func (r *MentorRepo) CreateMentor(mentor mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error) {
	var createdMentor mentor_dto.MentorJoinUserDto

	err := r.db.Db.QueryRow(
		`INSERT INTO mentors (user_id) VALUES ($1) RETURNING mentor_id, user_id`,
		mentor.UserId,
	).Scan(&createdMentor.MentorId, &createdMentor.Id)

	if err != nil {
		return nil, err
	}

	return &createdMentor, nil
}

func (r *MentorRepo) GetMentorById(id uint64) (*mentor_dto.MentorJoinUserDto, error) {
	var mentor mentor_dto.MentorJoinUserDto

	err := r.db.Db.QueryRow(
		`SELECT mentors.mentor_id, users.user_id, users.user_fullname, users.user_role
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE mentors.mentor_id = $1`,
		id,
	).Scan(&mentor.MentorId, &mentor.Id, &mentor.FullName, &mentor.Role)

	if err != nil {
		return nil, err
	}

	return &mentor, nil
}

func (r *MentorRepo) GetMentorByFullname(fullname string) (*mentor_dto.MentorJoinUserDto, error) {
	var mentor mentor_dto.MentorJoinUserDto

	err := r.db.Db.QueryRow(
		`SELECT mentors.mentor_id, users.user_id, users.user_fullname, users.user_role
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE users.user_fullname = $1`,
		fullname,
	).Scan(&mentor.MentorId, &mentor.Id, &mentor.FullName, &mentor.Role)

	if err != nil {
		return nil, err
	}

	return &mentor, nil
}
