package models

type Message struct {
	Id             uint64 `json:"message_id"`
	PublishingDate string `json:"publishing_date"`
	Type           string `json:"message_type"`
	DebtorId       uint64 `json:"debtor_id"`
	Author         string `json:"message_author"`
}
