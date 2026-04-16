package user_dto

type CreateUserDto struct {
	Fullname string `json:"user_fullname"`
	Role     string `json:"user_role"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
