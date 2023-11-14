package api

import (
	"net/http"

	"github.com/ca-lee-b/golander/server/internal/repositories"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func newUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) GetOneById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
