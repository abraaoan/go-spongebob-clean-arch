package postgres

import (
	"database/sql"

	"github.com/abraaoan/go-spongebob-clean-arch/internal/entity"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/repository"
)

type characterPostgres struct {
	db *sql.DB
}

func NewCharacterPostgres(db *sql.DB) repository.CharacterRepository {
	return &characterPostgres{db: db}
}

func (r *characterPostgres) Save(c *entity.Character) (string, error) {
	query := `INSERT INTO characters (id, name, description, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, c.ID, c.Name, c.Description, c.CreatedAt)
	return c.ID, err
}

func (r *characterPostgres) GetByID(id string) (*entity.Character, error) {
	var c entity.Character
	query := `SELECT id, name, description, created_at FROM characters WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *characterPostgres) List() ([]*entity.Character, error) {
	query := `SELECT id, name, description, created_at FROM characters ORDER BY name`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*entity.Character
	for rows.Next() {
		var c entity.Character
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt); err != nil {
			return nil, err
		}
		characters = append(characters, &c)
	}

	return characters, nil
}
