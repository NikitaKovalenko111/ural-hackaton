package requests_dto

type CreateRequestDto struct {
	Message string `json:"request_message"`
	UserId  uint64 `json:"user_id"`
}
