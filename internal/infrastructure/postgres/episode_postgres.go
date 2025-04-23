package postgres

import (
	"database/sql"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
)

type episodePostgres struct {
	db *sql.DB
}

func NewEpisodePostgres(db *sql.DB) repository.EpisodeRepository {
	return &episodePostgres{db: db}
}

func (r *episodePostgres) GetById(id string) (*entity.Episode, error) {
	var e entity.Episode
	query := `SELECT id, title, description, season_number, episode_number, air_date, created_at FROM episodes WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&e.ID, &e.Title, &e.Description, &e.SeasonNumber, &e.EpisodeNumber, &e.AirDate, &e.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *episodePostgres) List() ([]*entity.Episode, error) {
	query := `SELECT id, title, description, season_number, episode_number, air_date, created_at FROM episodes ORDER BY season_number, episode_number`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var episodes []*entity.Episode
	for rows.Next() {
		var e entity.Episode
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.SeasonNumber, &e.EpisodeNumber, &e.AirDate, &e.CreatedAt); err != nil {
			return nil, err
		}
		episodes = append(episodes, &e)
	}

	return episodes, nil
}

func (r *episodePostgres) ListBySeason(seasonNumber int) ([]*entity.Episode, error) {
	query := `SELECT id, title, description, season_number, episode_number, air_date, created_at FROM episodes WHERE season_number = $1 ORDER BY season_number, episode_number`
	rows, err := r.db.Query(query, seasonNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var episodes []*entity.Episode
	for rows.Next() {
		var e entity.Episode
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.SeasonNumber, &e.EpisodeNumber, &e.AirDate, &e.CreatedAt); err != nil {
			return nil, err
		}
		episodes = append(episodes, &e)
	}

	return episodes, nil
}

func (r *episodePostgres) Save(e *entity.Episode) (string, error) {
	query := `INSERT INTO episodes (id, title, description, season_number, episode_number, air_date, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Query(query, e.ID, e.Title, e.Description, e.SeasonNumber, e.EpisodeNumber, e.AirDate, e.CreatedAt)

	return e.ID, err
}
