package user_dto

type CreateUserDto struct {
	Fullname string `json:"fullname"`
	Role     string `json:"user_role"`
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
}
