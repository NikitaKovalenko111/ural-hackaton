package hubs_storage

import (
	hubsStorageDto "ural-hackaton/internal/dto/hub"
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
)

type HubRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *HubRepo {
	return &HubRepo{
		db: db,
	}
}

func (r *HubRepo) GetAllHubs() ([]models.Hub, error) {
	rows, err := r.db.Db.Query(
		`SELECT hub_id, hub_name, address, status, city, description, schedule, occupancy FROM hubs`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var hubs []models.Hub

	for rows.Next() {
		var hub models.Hub

		err = rows.Scan(&hub.Id, &hub.HubName, &hub.Address, &hub.Status, &hub.City, &hub.Description, &hub.Schedule, &hub.Occupancy)

		if err != nil {
			return nil, err
		}

		hubs = append(hubs, hub)
	}

	return hubs, nil
}

func (r *HubRepo) GetHubById(id uint64) (*models.Hub, error) {
	var hub models.Hub

	err := r.db.Db.QueryRow(
		`SELECT hub_id, hub_name, address, status, city, description, schedule, occupancy FROM hubs WHERE hub_id = $1`,
		id,
	).Scan(&hub.Id, &hub.HubName, &hub.Address, &hub.Status, &hub.City, &hub.Description, &hub.Schedule, &hub.Occupancy)

	if err != nil {
		return nil, err
	}

	return &hub, nil
}

func (r *HubRepo) CreateHub(hub *hubsStorageDto.CreateHubDto) (*models.Hub, error) {
	var createdHub models.Hub

	err := r.db.Db.QueryRow(
		`INSERT INTO hubs (hub_name, address, status, city, description, schedule, occupancy)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING hub_id, hub_name, address, status, city, description, schedule, occupancy`,
		hub.Name,
		hub.Address,
		hub.Status,
		hub.City,
		hub.Description,
		hub.Schedule,
		hub.Occupancy,
	).Scan(&createdHub.Id, &createdHub.HubName, &createdHub.Address, &createdHub.Status, &createdHub.City, &createdHub.Description, &createdHub.Schedule, &createdHub.Occupancy)

	if err != nil {
		return nil, err
	}

	return &createdHub, nil
}

func (r *HubRepo) UpdateHub(hub *models.Hub) (*models.Hub, error) {
	var updatedHub models.Hub

	err := r.db.Db.QueryRow(
		`UPDATE hubs
		 SET hub_name = $1, address = $2, status = $3, city = $4, description = $5, schedule = $6, occupancy = $7
		 WHERE hub_id = $8
		 RETURNING hub_id, hub_name, address, status, city, description, schedule, occupancy`,
		hub.HubName, hub.Address, hub.Status, hub.City, hub.Description, hub.Schedule, hub.Occupancy, hub.Id,
	).Scan(&updatedHub.Id, &updatedHub.HubName, &updatedHub.Address, &updatedHub.Status, &updatedHub.City, &updatedHub.Description, &updatedHub.Schedule, &updatedHub.Occupancy)

	if err != nil {
		return nil, err
	}

	return &updatedHub, nil
}

func (r *HubRepo) DeleteHub(id uint64) error {
	_, err := r.db.Db.Exec(
		`DELETE FROM hubs WHERE hub_id = $1`,
		id,
	)

	return err
}
