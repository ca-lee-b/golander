package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ca-lee-b/golander/server/internal/models"
	"github.com/lib/pq"
)

type EventRepository struct {
	db *sql.DB
}

func newEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) AddParticipant(id int, participant string) error {
	_, err := r.db.Exec("UPDATE events SET participants = array_append(participants, $1) WHERE id = $2", participant, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) GetOneById(id int) *models.Event {
	var event models.Event
	row := r.db.QueryRow("SELECT * FROM events WHERE id = $1", id)

	err := row.Scan(&event.Id, &event.Title, &event.Owner_id, &event.Created_at, pq.Array(&event.Available_dates), &event.Participants)
	if err != nil {
		fmt.Printf("err %v", err)
		return nil
	}

	return &event
}

func (r *EventRepository) CreateEvent(title string, ownerId string, availableDates string) error {
	_, err := r.db.Exec("INSERT INTO events(title, owner_id, available_dates) VALUES($1, $2, $3)", title, ownerId, pq.Array(availableDates))
	if err != nil {
		return err
	}

	return nil
}
