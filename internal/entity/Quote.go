package entity

import "time"

type Quote struct {
	ID          string
	Text        string
	CharacterID string
	EpisodeID   *string
	CreatedAt   time.Time
}
