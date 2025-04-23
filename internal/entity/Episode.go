package entity

import "time"

type Episode struct {
	ID            string
	Title         string
	Description   string
	SeasonNumber  int
	EpisodeNumber int
	AirDate       time.Time
	CreatedAt     time.Time
}
