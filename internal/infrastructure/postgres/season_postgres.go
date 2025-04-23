package postgres

import (
	"database/sql"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
)

type seasonPostgres struct {
	db *sql.DB
}

func NewSeasonPostgres(db *sql.DB) repository.SeasonRepository {
	return &seasonPostgres{db: db}
}

func (r *seasonPostgres) GetById(id string) (*entity.Season, error) {
	var s entity.Season
	query := `SELECT id, number, description, created_at FROM seasons WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&s.ID, &s.Number, &s.Description, &s.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *seasonPostgres) List() ([]*entity.Season, error) {
	query := `SELECT id, number, description, created_at FROM seasons ORDER BY number`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seasons []*entity.Season
	for rows.Next() {
		var s entity.Season
		if err := rows.Scan(&s.ID, &s.Number, &s.Description, &s.CreatedAt); err != nil {
			return nil, err
		}
		seasons = append(seasons, &s)
	}

	return seasons, nil
}

func (r *seasonPostgres) Save(s *entity.Season) (string, error) {
	query := `INSERT INTO seasons (id, number, description, created_at) VALUES ($1, $2, $3, $3)`
	_, err := r.db.Exec(query, s.ID, s.Number, s.Description, s.CreatedAt)

	return s.ID, err
}
