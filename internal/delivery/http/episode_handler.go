package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/usecase"
)

type EpisodeHandler struct {
	UC usecase.EpisodeUseCase
}

func NewEpisodeHandler(uc usecase.EpisodeUseCase) *EpisodeHandler {
	return &EpisodeHandler{UC: uc}
}

func (h *EpisodeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/episodes/")
	if id == "" {
		http.Error(w, "missing episode ID", http.StatusBadRequest)
		return
	}

	episode, err := h.UC.GetById(id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if episode == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episode)
}

func (h *EpisodeHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var ep entity.Episode
	if err := json.NewDecoder(r.Body).Decode(&ep); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.UC.Create(&ep)
	if err != nil {
		http.Error(w, "failed to create episode", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
