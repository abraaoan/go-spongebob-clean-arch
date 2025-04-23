package postgres

import (
	"database/sql"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
)

type quotePostgres struct {
	db *sql.DB
}

func NewQuotePostgres(db *sql.DB) repository.QuoteRepository {
	return &quotePostgres{db: db}
}

func (r *quotePostgres) Save(q *entity.Quote) (string, error) {
	query := `INSERT INTO quotes (id, text, character_id, episode_id, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, q.ID, q.Text, q.CharacterID, q.EpisodeID, q.CreatedAt)
	return q.ID, err
}

func (r *quotePostgres) GetByID(id string) (*entity.Quote, error) {
	var q entity.Quote
	query := `SELECT id, text, character_id, episode_id, created_at FROM quotes WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&q.ID, &q.Text, &q.CharacterID, &q.EpisodeID, &q.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &q, nil
}

func (r *quotePostgres) ListByCharacter(characterID string) ([]*entity.Quote, error) {
	query := `SELECT id, text, character_id, episode_id, created_at FROM quotes WHERE character_id = $1`
	rows, err := r.db.Query(query, characterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*entity.Quote
	for rows.Next() {
		var q entity.Quote
		if err := rows.Scan(&q.ID, &q.Text, &q.CharacterID, &q.EpisodeID, &q.CreatedAt); err != nil {
			return nil, err
		}
		quotes = append(quotes, &q)
	}
	return quotes, nil
}

func (r *quotePostgres) ListByEpisode(episodeID string) ([]*entity.Quote, error) {
	query := `SELECT id, text, character_id, episode_id, created_at FROM quotes WHERE episode_id = $1`
	rows, err := r.db.Query(query, episodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*entity.Quote
	for rows.Next() {
		var q entity.Quote
		if err := rows.Scan(&q.ID, &q.Text, &q.CharacterID, &q.EpisodeID, &q.CreatedAt); err != nil {
			return nil, err
		}
		quotes = append(quotes, &q)
	}
	return quotes, nil
}
