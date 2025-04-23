package usecase

import (
	"errors"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
	"github.com/google/uuid"
)

type QuoteUseCase interface {
	Create(quote *entity.Quote) (string, error)
	GetByID(id string) (*entity.Quote, error)
	ListByCharacter(characterID string) ([]*entity.Quote, error)
	ListByEpisode(episodeID string) ([]*entity.Quote, error)
}

type quoteUseCase struct {
	repo repository.QuoteRepository
}

func NewQuoteUseCase(repo repository.QuoteRepository) QuoteUseCase {
	return &quoteUseCase{repo: repo}
}

func (uc *quoteUseCase) Create(quote *entity.Quote) (string, error) {
	if quote.Text == "" || quote.CharacterID == "" {
		return "", errors.New("quote text and character ID are required")
	}

	quote.ID = uuid.New().String()
	quote.CreatedAt = time.Now()

	return uc.repo.Save(quote)
}

func (uc *quoteUseCase) GetByID(id string) (*entity.Quote, error) {
	return uc.repo.GetByID(id)
}

func (uc *quoteUseCase) ListByCharacter(characterID string) ([]*entity.Quote, error) {
	return uc.repo.ListByCharacter(characterID)
}

func (uc *quoteUseCase) ListByEpisode(episodeID string) ([]*entity.Quote, error) {
	return uc.repo.ListByEpisode(episodeID)
}
