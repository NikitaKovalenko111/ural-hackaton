package user_dto

type CreateUserDto struct {
	Fullname string `json:"user_fullname"`
	Role     string `json:"user_role"`
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
	Phone    string `json:"phone"`
}
