package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sh0e1/wire/domain"
	"github.com/sh0e1/wire/usecase"
)

func NewTodoHandler(interactor *usecase.TodoInteractor) *TodoHandler {
	return &TodoHandler{
		interactor: interactor,
	}
}

type TodoHandler struct {
	interactor *usecase.TodoInteractor
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.interactor.Create(r.Context(), &todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todo, err := h.interactor.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	todos, err := h.interactor.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
