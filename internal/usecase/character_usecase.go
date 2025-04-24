package usecase

import (
	"errors"
	"time"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/infrastructure/cache"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
	"github.com/google/uuid"
)

type CharacterUseCase interface {
	Create(character *entity.Character) (string, error)
	GetByID(id string) (*entity.Character, error)
	List() ([]*entity.Character, error)
}

type characterUseCase struct {
	repo  repository.CharacterRepository
	cache *cache.CharacterCache
}

func NewCharacterUseCase(repo repository.CharacterRepository, cache *cache.CharacterCache) CharacterUseCase {
	return &characterUseCase{repo: repo, cache: cache}
}

func (uc *characterUseCase) Create(character *entity.Character) (string, error) {
	if character.Name == "" {
		return "", errors.New("character name is required")
	}

	character.ID = uuid.New().String()
	character.CreatedAt = time.Now()

	return uc.repo.Save(character)
}

func (uc *characterUseCase) GetByID(id string) (*entity.Character, error) {
	if cached, ok := uc.cache.Get(id); ok {
		return cached, nil
	}

	character, err := uc.repo.GetByID(id)
	if err != nil || character == nil {
		return character, err
	}

	// Save to cache.
	uc.cache.Set(id, character)

	return character, nil
}

func (uc *characterUseCase) List() ([]*entity.Character, error) {
	return uc.repo.List()
}
