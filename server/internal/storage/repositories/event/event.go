package event_storage

import (
	"database/sql"
	"fmt"
	event_dto "ural-hackaton/internal/dto/event"
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
)

type EventRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (r *EventRepo) CreateEvent(event *event_dto.CreateEventDto) error {
	_, err := r.db.Db.Exec(
		`INSERT INTO events (name, description, start_time, end_time, hub_id) VALUES ($1, $2, $3, $4, $5)`,
		event.Name,
		event.Description,
		event.StartTime,
		event.EndTime,
		event.HubId,
	)

	if err != nil {
		return fmt.Errorf("Couldn't create event!")
	}

	return nil
}

func (r *EventRepo) GetEventById(id uint64) (*models.Event, error) {
	var event models.Event

	err := r.db.Db.QueryRow(
		`SELECT event_id, name, description, start_time, end_time, hub_id FROM events WHERE event_id = $1`,
		id,
	).Scan(&event.Id, &event.EventName, &event.Description, &event.StartTime, &event.EndTime, &event.HubId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with this id not found!")
		}

		return nil, fmt.Errorf("Couldn't get event by id!")
	}

	return &event, nil
}

func (r *EventRepo) GetEventByName(name string) (*models.Event, error) {
	var event models.Event

	err := r.db.Db.QueryRow(
		`SELECT event_id, name, description, start_time, end_time, hub_id FROM events WHERE name = $1`,
		name,
	).Scan(&event.Id, &event.EventName, &event.Description, &event.StartTime, &event.EndTime, &event.HubId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with this name not found!")
		}

		return nil, fmt.Errorf("Couldn't get event by name!")
	}

	return &event, nil
}

func (r *EventRepo) GetEventsByHubId(hubId uint64) ([]models.Event, error) {
	var events []models.Event

	rows, err := r.db.Db.Query(
		`SELECT event_id, name, description, start_time, end_time, hub_id FROM events WHERE hub_id = $1`,
		hubId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with this user id not found!")
		}

		return nil, fmt.Errorf("Couldn't get event by user id!")
	}

	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.Id, &event.EventName, &event.Description, &event.StartTime, &event.EndTime, &event.HubId); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepo) GetAllEvents() ([]models.Event, error) {
	rows, err := r.db.Db.Query(
		`SELECT event_id, name, description, start_time, end_time, hub_id FROM events`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event

		err = rows.Scan(&event.Id, &event.EventName, &event.Description, &event.StartTime, &event.EndTime, &event.HubId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepo) UpdateEvent(event *models.Event) (*models.Event, error) {
	var updatedEvent models.Event

	err := r.db.Db.QueryRow(
		`UPDATE events SET name = $1, description = $2, start_time = $3, end_time = $4, hub_id = $5 WHERE event_id = $6 RETURNING event_id, name, description, start, end, hub_id`,
		event.EventName, event.Description, event.StartTime, event.EndTime, event.HubId, event.Id,
	).Scan(&updatedEvent.Id, &updatedEvent.EventName, &updatedEvent.Description, &updatedEvent.StartTime, &updatedEvent.EndTime, &updatedEvent.HubId)

	if err != nil {
		return nil, err
	}

	return &updatedEvent, nil
}

func (r *EventRepo) DeleteEvent(id uint64) error {
	_, err := r.db.Db.Exec(
		`DELETE FROM events WHERE event_id = $1`,
		id,
	)

	return err
}
