package posts

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.store.List())
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	post := h.store.Create(input.Title, input.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	post, ok := h.store.Get(id)
	if !ok {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
