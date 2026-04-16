package hubs_storage

import (
	hubsStorageDto "ural-hackaton/internal/dto/hubs"
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
)

type HubsRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *HubsRepo {
	return &HubsRepo{
		db: db,
	}
}

func (r *HubsRepo) GetAllHubs() ([]*models.Hub, error) {
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

func (r *HubsRepo) GetHubById(id uint64) (*models.Hub, error) {
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

func (r *HubsRepo) CreateHub(hub *hubsStorageDto.CreateHubDto) (*models.Hub, error) {
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

func (r *HubsRepo) UpdateHub(hub *models.Hub) (*models.Hub, error) {
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

func (r *HubsRepo) DeleteHub(id uint64) error {
	_, err := r.db.Db.Exec(
		`DELETE FROM hubs WHERE hub_id = $1`,
		id,
	)

	return err
}
