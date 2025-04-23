package repository

import "github.com/abraaoan/go-spongebob-clean-arch/internal/entity"

type QuoteRepository interface {
	Save(quote *entity.Quote) (string, error)
	GetByID(id string) (*entity.Quote, error)
	ListByCharacter(characterID string) ([]*entity.Quote, error)
	ListByEpisode(episodeID string) ([]*entity.Quote, error)
}
