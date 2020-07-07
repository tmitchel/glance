package definition

import "github.com/google/uuid"

type CardService interface {
	New(CreateRequest) CardResponse
	Card(CardRequest) CardResponse
	Cards(CardsRequest) CardsResponse
}

type CreateRequest struct {
	Title   string
	Content string
	Creator string
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
