package hubs_storage

import (
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
	hubsStorageDto "ural-hackaton/internal/storage/repositories/hubs/dto"
)

type HubsRepository struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *HubsRepository {
	return &HubsRepository{
		db: db,
	}
}

func (r *HubsRepository) GetAllHubs() ([]*models.Hub, error) {
	rows, err := r.db.Db.Query(
		`SELECT hub_id, hub_name FROM hubs`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var hubs []*models.Hub

	for rows.Next() {
		var hub models.Hub

		err = rows.Scan(&hub.Id, &hub.HubName)

		if err != nil {
			return nil, err
		}

		hubs = append(hubs, &hub)
	}

	return hubs, nil
}

func (r *HubsRepository) GetHubById(id uint64) (*models.Hub, error) {
	var hub models.Hub

	err := r.db.Db.QueryRow(
		`SELECT hub_id, hub_name FROM hubs WHERE hub_id = $1`,
		id,
	).Scan(&hub.Id, &hub.HubName)

	if err != nil {
		return nil, err
	}

	return &hub, nil
}

func (r *HubsRepository) CreateHub(hub *hubsStorageDto.CreateHubDto) (*models.Hub, error) {
	var createdHub models.Hub

	err := r.db.Db.QueryRow(
		`INSERT INTO hubs (hub_name) VALUES ($1) RETURNING hub_id, hub_name`,
		hub.Name,
	).Scan(&createdHub.Id, &createdHub.HubName)

	if err != nil {
		return nil, err
	}

	return &createdHub, nil
}

func (r *HubsRepository) UpdateHub(hub *models.Hub) (*models.Hub, error) {
	var updatedHub models.Hub

	err := r.db.Db.QueryRow(
		`UPDATE hubs SET hub_name = $1 WHERE hub_id = $2 RETURNING hub_id, hub_name`,
		hub.HubName, hub.Id,
	).Scan(&updatedHub.Id, &updatedHub.HubName)

	if err != nil {
		return nil, err
	}

	return &updatedHub, nil
}

func (r *HubsRepository) DeleteHub(id uint64) error {
	_, err := r.db.Db.Exec(
		`DELETE FROM hubs WHERE hub_id = $1`,
		id,
	)

	return err
}
