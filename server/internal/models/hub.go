package models

type Hub struct {
	Id      uint64 `json:"hub_id"`
	HubName string `json:"hub_name"`
	Address string `json:"address"`
	Status  string `json:"status"`
}
