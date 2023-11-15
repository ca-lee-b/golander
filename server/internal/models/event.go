package models

import "time"

type Event struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Owner_id        string    `json:"owner_id"`
	Created_at      time.Time `json:"created_at"`
	Available_dates string    `json:"available_dates"`
}
