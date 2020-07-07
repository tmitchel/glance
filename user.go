package glance

import "github.com/google/uuid"

// User represents a single user of the app. It is very
// minimal at the moment.
type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}
