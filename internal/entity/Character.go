package entity

import "time"

type Character struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}
