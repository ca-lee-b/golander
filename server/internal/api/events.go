package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ca-lee-b/golander/server/internal/repositories"
	"github.com/go-chi/chi/v5"
)

type EventHandler struct {
	eventRepository *repositories.EventRepository
}

func newEventHandler(repo *repositories.EventRepository) *EventHandler {
	return &EventHandler{
		eventRepository: repo,
	}
}

func (h *EventHandler) AddParticipant(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	if param == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	event := h.eventRepository.GetOneById(id)
	if event == nil {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	type bodyData struct {
		Id string `json:"id"`
	}
	var data bodyData

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Printf("error: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = h.eventRepository.AddParticipant(event.Id, data.Id)
	if err != nil {
		fmt.Printf("error: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Success"))
}

func (h *EventHandler) GetOneById(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	if param == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	event := h.eventRepository.GetOneById(id)
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

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	type createEvent struct {
		Title           string `json:"title"`
		Owner_id        string `json:"owner_id"`
		Available_dates string `json:"available_dates"`
	}
	var data createEvent

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	if data.Title == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	if data.Owner_id == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	json, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(200)
	w.Write(json)
}
