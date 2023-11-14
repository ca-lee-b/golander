package repositories

import (
	"database/sql"
)

type Repositories struct {
	UserRepository *UserRepository
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: newUserRepository(db),
	}
}
