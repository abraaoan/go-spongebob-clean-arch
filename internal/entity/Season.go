package entity

import "time"

type Season struct {
	ID          string
	Number      int
	Description string
	CreatedAt   time.Time
}
