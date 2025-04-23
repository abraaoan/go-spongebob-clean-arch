package repository

import "github.com/abraaoan/go-spongebob-clean-arch/internal/entity"

type SeasonRepository interface {
	Save(season *entity.Season) (string, error)
	GetById(id string) (*entity.Season, error)
	List() ([]*entity.Season, error)
}
