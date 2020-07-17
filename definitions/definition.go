package definition

import "github.com/google/uuid"

type CreateService interface {
	Card(CreateCardRequest) CardResponse
	User(CreateUserRequest) UserResponse
	ClaimCard(ClaimRequest) CardResponse
	UpdateStatus(NewStatusRequest) CardResponse
	Finalize(ClaimRequest) CardResponse
}

type GetService interface {
	Card(CardRequest) CardResponse
	Cards(EmptyRequest) CardsResponse
	User(UserRequest) UserResponse
	Users(EmptyRequest) UsersResponse
	HomePage(UserRequest) HomePageResponse
}

type CreateCardRequest struct {
	Title   string
	Content string
	Creator string
}

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type ClaimRequest struct {
	UserID string
	CardID string
}

type NewStatusRequest struct {
	Status int
	Card   string
	User   string
}

type CardResponse struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Status    int
	Creator   uuid.UUID
	Claimed   bool
	CreatedAt string
}

type UserResponse struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type CardRequest struct {
	ID string
}

type UserRequest struct {
	Email string
}

type EmptyRequest struct{}

type UsersResponse struct {
	Users []UserResponse
}

type CardsResponse struct {
	Cards []CardResponse
}

type HomePageResponse struct {
	User  *UserResponse
	Card  *CardResponse
	Pairs []struct {
		User *UserResponse
		Card *CardResponse
	}
}
