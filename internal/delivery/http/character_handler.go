package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/usecase"
)

type CharacterHandler struct {
	UC usecase.CharacterUseCase
}

func NewCharacterHandler(uc usecase.CharacterUseCase) *CharacterHandler {
	return &CharacterHandler{UC: uc}
}

func (h *CharacterHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/characters/")

	if id == "" {
		http.Error(w, "missing charactter ID", http.StatusBadRequest)
	}

	character, err := h.UC.GetByID(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}

func (h *CharacterHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var character entity.Character
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.UC.Create(&character)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create character %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
