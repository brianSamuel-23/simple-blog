package entity

import "time"

// Timestamps holds the lifecycle fields shared by entities that support updates.
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
