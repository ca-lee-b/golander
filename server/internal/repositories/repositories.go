package repositories

import (
	"database/sql"
)

type Repositories struct {
	UserRepository  *UserRepository
	EventRepository *EventRepository
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository:  newUserRepository(db),
		EventRepository: newEventRepository(db),
	}
}
