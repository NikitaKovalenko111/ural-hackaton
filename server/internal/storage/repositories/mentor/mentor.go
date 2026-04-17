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

func nullableHubID(value sql.NullInt64) *uint64 {
	if !value.Valid {
		return nil
	}

	hubID := uint64(value.Int64)
	return &hubID
}

func (r *MentorRepo) CreateMentor(mentor mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error) {
	var createdMentor mentor_dto.MentorJoinUserDto
	var hubID sql.NullInt64

	err := r.db.Db.QueryRow(
		`INSERT INTO mentors (user_id, hub_id) VALUES ($1, $2) RETURNING mentor_id, user_id, hub_id`,
		mentor.UserId,
		mentor.HubId,
	).Scan(&createdMentor.MentorId, &createdMentor.Id, &hubID)

	if err != nil {
		return nil, err
	}

	createdMentor.HubId = nullableHubID(hubID)

	return &createdMentor, nil
}

func (r *MentorRepo) GetMentorById(id uint64) (*mentor_dto.MentorJoinUserDto, error) {
	var mentor mentor_dto.MentorJoinUserDto
	var hubID sql.NullInt64

	err := r.db.Db.QueryRow(
		`SELECT mentors.mentor_id, mentors.hub_id, users.user_id, users.fullname, users.user_role, users.email, users.telegram, users.phone
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE mentors.mentor_id = $1`,
		id,
	).Scan(&mentor.MentorId, &hubID, &mentor.Id, &mentor.FullName, &mentor.Role, &mentor.Email, &mentor.Telegram, &mentor.Phone)

	if err != nil {
		return nil, err
	}

	mentor.HubId = nullableHubID(hubID)

	return &mentor, nil
}

func (r *MentorRepo) GetMentorByFullname(fullname string) (*mentor_dto.MentorJoinUserDto, error) {
	var mentor mentor_dto.MentorJoinUserDto
	var hubID sql.NullInt64

	err := r.db.Db.QueryRow(
		`SELECT mentors.mentor_id, mentors.hub_id, users.user_id, users.fullname, users.user_role, users.email, users.telegram, users.phone
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE users.fullname = $1`,
		fullname,
	).Scan(&mentor.MentorId, &hubID, &mentor.Id, &mentor.FullName, &mentor.Role, &mentor.Email, &mentor.Telegram, &mentor.Phone)

	if err != nil {
		return nil, err
	}

	mentor.HubId = nullableHubID(hubID)

	return &mentor, nil
}

func (r *MentorRepo) GetMentorByUserId(userId uint64) (*mentor_dto.MentorJoinUserDto, error) {
	var mentor mentor_dto.MentorJoinUserDto
	var hubID sql.NullInt64

	err := r.db.Db.QueryRow(
		`SELECT mentors.mentor_id, mentors.hub_id, users.user_id, users.fullname, users.user_role, users.email, users.telegram, users.phone
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 WHERE mentors.user_id = $1`,
		userId,
	).Scan(&mentor.MentorId, &hubID, &mentor.Id, &mentor.FullName, &mentor.Role, &mentor.Email, &mentor.Telegram, &mentor.Phone)

	if err != nil {
		return nil, err
	}

	mentor.HubId = nullableHubID(hubID)

	return &mentor, nil
}

func (r *MentorRepo) GetAllMentors() ([]mentor_dto.MentorJoinUserDto, error) {
	rows, err := r.db.Db.Query(
		`SELECT mentors.mentor_id, mentors.hub_id, users.user_id, users.fullname, users.user_role, users.email, users.telegram, users.phone
		 FROM mentors
		 JOIN users ON mentors.user_id = users.user_id
		 ORDER BY users.fullname`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mentors := make([]mentor_dto.MentorJoinUserDto, 0)
	for rows.Next() {
		var mentor mentor_dto.MentorJoinUserDto
		var hubID sql.NullInt64
		if err = rows.Scan(&mentor.MentorId, &hubID, &mentor.Id, &mentor.FullName, &mentor.Role, &mentor.Email, &mentor.Telegram, &mentor.Phone); err != nil {
			return nil, err
		}

		mentor.HubId = nullableHubID(hubID)

		mentors = append(mentors, mentor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return mentors, nil
}
