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
		`SELECT hub_id, hub_name, address, status FROM hubs`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var hubs []models.Hub

	for rows.Next() {
		var hub models.Hub

		err = rows.Scan(&hub.Id, &hub.HubName, &hub.Address, &hub.Status)

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
		`SELECT hub_id, hub_name, address, status FROM hubs WHERE hub_id = $1`,
		id,
	).Scan(&hub.Id, &hub.HubName, &hub.Address, &hub.Status)

	if err != nil {
		return nil, err
	}

	return &hub, nil
}

func (r *HubRepo) CreateHub(hub *hubsStorageDto.CreateHubDto) (*models.Hub, error) {
	var createdHub models.Hub

	err := r.db.Db.QueryRow(
		`INSERT INTO hubs (hub_name, address, status) VALUES ($1, $2, $3) RETURNING hub_id, hub_name, address, status`,
		hub.Name,
		hub.Address,
		"opened",
	).Scan(&createdHub.Id, &createdHub.HubName, &createdHub.Address, &createdHub.Status)

	if err != nil {
		return nil, err
	}

	return &createdHub, nil
}

func (r *HubRepo) UpdateHub(hub *models.Hub) (*models.Hub, error) {
	var updatedHub models.Hub

	err := r.db.Db.QueryRow(
		`UPDATE hubs SET hub_name = $1, address = $2, status = $3 WHERE hub_id = $4 RETURNING hub_id, hub_name, address, status`,
		hub.HubName, hub.Address, hub.Status, hub.Id,
	).Scan(&updatedHub.Id, &updatedHub.HubName, &updatedHub.Address, &updatedHub.Status)

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
