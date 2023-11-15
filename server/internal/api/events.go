package api

import (
	"encoding/json"
	"net/http"

	"github.com/ca-lee-b/golander/server/internal/repositories"
)

type EventHandler struct {
	eventRepository *repositories.EventRepository
}

func newEventHandler(repo *repositories.EventRepository) *EventHandler {
	return &EventHandler{
		eventRepository: repo,
	}
}

func (h *EventHandler) GetOneById(w http.ResponseWriter, r *http.Request) {
	event := h.eventRepository.GetOneById(1)
	if event == nil {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	data, err := json.Marshal(&event)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
	}

	w.WriteHeader(200)
	w.Write(data)

}
