package models

type Requests struct {
	Id       uint64  `json:"request_id"`
	Message  string  `json:"request_message"`
	UserId   uint64  `json:"user_id"`
	MentorId *uint64 `json:"mentor_id,omitempty"`
}
