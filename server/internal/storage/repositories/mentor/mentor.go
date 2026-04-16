package mentor_storage

import (
	"database/sql"

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

func (r *MentorRepo) GetMentorsByFullname(fullname string) ([]*mentor_dto.MentorJoinUserDto, error) {
	rows, err := r.db.Db.Query(
		`SELECT mentors.mentor_id, users.user_id, users.user_fullname, users.user_role
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE users.user_fullname = $1`,
		fullname,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mentors []*mentor_dto.MentorJoinUserDto
	for rows.Next() {
		var mentor mentor_dto.MentorJoinUserDto
		err := rows.Scan(&mentor.MentorId, &mentor.Id, &mentor.FullName, &mentor.Role)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, &mentor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(mentors) == 0 {
		return nil, sql.ErrNoRows
	}

	return mentors, nil
}

func (r *MentorRepo) GetMentorsByRole(role string) ([]*mentor_dto.MentorJoinUserDto, error) {
	rows, err := r.db.Db.Query(
		`SELECT mentors.mentor_id, users.user_id, users.user_fullname, users.user_role
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE users.user_role = $1`,
		role,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mentors []*mentor_dto.MentorJoinUserDto
	for rows.Next() {
		var mentor mentor_dto.MentorJoinUserDto
		err := rows.Scan(&mentor.MentorId, &mentor.Id, &mentor.FullName, &mentor.Role)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, &mentor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(mentors) == 0 {
		return nil, sql.ErrNoRows
	}

	return mentors, nil
}
