package health

import (
	"errors"
	"net/http"

	"github.com/shashank/dynamodb-go-crud/internal/handlers"
	"github.com/shashank/dynamodb-go-crud/internal/repository/adapter"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func Get(h *Handler) func(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}
	HttpStatus.StatusOk(w, r, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}
func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}
func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
