package repository

import "github.com/abraaoan/go-spongebob-clean-arch/internal/entity"

type CharacterRepository interface {
	Save(character *entity.Character) (string, error)
	GetByID(id string) (*entity.Character, error)
	List() ([]*entity.Character, error)
}
