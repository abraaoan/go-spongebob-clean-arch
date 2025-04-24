package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/usecase"
)

type SeasonHandler struct {
	UC usecase.SeasonUseCase
}

func NewSeasonHandler(uc usecase.SeasonUseCase) *SeasonHandler {
	return &SeasonHandler{UC: uc}
}

func (h *SeasonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/seasons/")
	if id == "" {
		http.Error(w, "missing season ID", http.StatusBadRequest)
		return
	}

	season, err := h.UC.GetByID(id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if season == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(season)
}

func (h *SeasonHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var s entity.Season
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.UC.Create(&s)
	if err != nil {
		http.Error(w, "failed to create season", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
