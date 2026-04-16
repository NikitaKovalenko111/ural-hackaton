package models

import "database/sql"

type Admin struct {
	Id              uint64         `json:"admin_id"`
	Email           string         `json:"admin_email"`
	FullName        string         `json:"admin_fullname"`
	Password        sql.NullString `json:"admin_password"`
	ActivationToken string         `json:"activation_token"`
	IsActivated     bool           `json:"is_activated"`
}
