package usecase

import (
	"errors"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
	"github.com/google/uuid"
)

type SeasonUseCase interface {
	Create(season *entity.Season) (string, error)
	GetByID(id string) (*entity.Season, error)
	List() ([]*entity.Season, error)
}

type seasonUseCase struct {
	repo repository.SeasonRepository
}

func NewSeasonUseCase(repo repository.SeasonRepository) SeasonUseCase {
	return &seasonUseCase{repo: repo}
}

func (uc *seasonUseCase) Create(season *entity.Season) (string, error) {
	if season.Number <= 0 {
		return "", errors.New("invalid season number")
	}

	season.ID = uuid.New().String()
	season.CreatedAt = time.Now()

	return uc.repo.Save(season)
}

func (uc *seasonUseCase) GetByID(id string) (*entity.Season, error) {
	return uc.repo.GetById(id)
}

func (uc *seasonUseCase) List() ([]*entity.Season, error) {
	return uc.repo.List()
}
