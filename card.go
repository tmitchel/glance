package glance

import (
	"time"

	"github.com/google/uuid"
)

// Card represents a single task.
type Card struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Status    int
	Creator   uuid.UUID
	Volunteer uuid.UUID
	CreatedAt time.Time
}
