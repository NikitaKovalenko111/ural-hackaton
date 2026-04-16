package models

type User struct {
	Id       uint64 `json:"user_id"`
	FullName string `json:"user_fullname"`
	Role     string `json:"user_role"`
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
}
