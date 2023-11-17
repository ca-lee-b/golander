package models

import (
	"time"

	"github.com/lib/pq"
)

type Event struct {
	Id              int            `json:"id"`
	Title           string         `json:"title"`
	Owner_id        string         `json:"owner_id"`
	Created_at      time.Time      `json:"created_at"`
	Available_dates []time.Time    `json:"available_dates"`
	Participants    pq.StringArray `json:"participants"`
}
