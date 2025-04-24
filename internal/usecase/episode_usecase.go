package usecase

import (
	"errors"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/infrastructure/cache"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
	"github.com/google/uuid"
)

type EpisodeUseCase interface {
	Create(episode *entity.Episode) (string, error)
	GetById(id string) (*entity.Episode, error)
	List() ([]*entity.Episode, error)
	ListBySeason(seasonNumber int) ([]*entity.Episode, error)
}

type episodeUseCase struct {
	repo  repository.EpisodeRepository
	cache *cache.EpisodeCache
}

func NewEpisodeUseCase(repo repository.EpisodeRepository, cache *cache.EpisodeCache) EpisodeUseCase {
	return &episodeUseCase{repo: repo, cache: cache}
}

func (uc *episodeUseCase) Create(episode *entity.Episode) (string, error) {
	if episode.Title == "" {
		return "", errors.New("episode title is required")
	}

	episode.ID = uuid.New().String()
	episode.CreatedAt = time.Now()

	return uc.repo.Save(episode)
}

func (uc *episodeUseCase) GetById(id string) (*entity.Episode, error) {
	if e, ok := uc.cache.Get(id); ok {
		return e, nil
	}

	e, err := uc.repo.GetById(id)
	if err != nil || e == nil {
		return e, err
	}

	uc.cache.Set(id, e)
	return e, nil
}

func (uc *episodeUseCase) List() ([]*entity.Episode, error) {
	return uc.repo.List()
}

func (uc *episodeUseCase) ListBySeason(seasonNumber int) ([]*entity.Episode, error) {
	return uc.repo.ListBySeason(seasonNumber)
}
