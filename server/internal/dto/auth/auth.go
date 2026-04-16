package auth_dto

// RequestMagicLinkDto — входящие данные для запроса ссылки
type RequestMagicLinkDto struct {
	Email string `json:"email" validate:"required,email"`
}

// VerifyMagicLinkResponse — данные, возвращаемые после успешной верификации
// (используется внутри сервиса/контроллера для формирования ответа)
type VerifyMagicLinkResponse struct {
	UserID       uint64
	Fullname     string
	Email        string
	Role         string
	SessionToken string // или JWT
}
