package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/usecase"
)

type QuoteHandler struct {
	UC usecase.QuoteUseCase
}

func NewQuoteHandler(uc usecase.QuoteUseCase) *QuoteHandler {
	return &QuoteHandler{UC: uc}
}

func (h *QuoteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/quotes/")
	if id == "" {
		http.Error(w, "missing quote ID", http.StatusBadRequest)
		return
	}

	quote, err := h.UC.GetByID(id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if quote == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func (h *QuoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var q entity.Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.UC.Create(&q)
	if err != nil {
		http.Error(w, "failed to create quote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
