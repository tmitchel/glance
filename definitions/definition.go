package definition

import "github.com/google/uuid"

type CardService interface {
	Cards(CardsRequest) CardsResponse
	Card(CardRequest) CardResponse
}

type CardRequest struct {
	ID string
}

type CardResponse struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Status    int
	Creator   uuid.UUID
	Volunteer uuid.UUID
	CreatedAt string
}

type CardsRequest struct{}

type CardsResponse struct {
	Cards []CardResponse
}
