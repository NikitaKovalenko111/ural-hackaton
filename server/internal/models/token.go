package models

type Token struct {
	Id           uint   `json:"token_id"`
	RefreshToken string `json:"refresh_token"`
	AdminId      uint   `json:"admin_id"`
}
