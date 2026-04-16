package hubs_dto

type CreateHubDto struct {
	Name        string `json:"hub_name"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	City        string `json:"city"`
	Description string `json:"description"`
	Schedule    string `json:"schedule"`
	Occupancy   int    `json:"occupancy"`
}
