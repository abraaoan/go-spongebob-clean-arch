package repository

import "github.com/abraaoan/go-spongebob-clean-arch/internal/entity"

type EpisodeRepository interface {
	Save(episode *entity.Episode) (string, error)
	GetById(id string) (*entity.Episode, error)
	List() ([]*entity.Episode, error)
	ListBySeason(seasonNumber int) ([]*entity.Episode, error)
}
