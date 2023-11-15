package api

import (
	"encoding/json"
	"net/http"

	"github.com/ca-lee-b/golander/server/internal/repositories"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

func newUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: repo,
	}
}

func (h *UserHandler) GetOneById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	user := h.userRepository.GetOneById(id)
	if user == nil {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	data, err := json.Marshal(map[string]interface{}{
		"email": user.Email,
	})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}
