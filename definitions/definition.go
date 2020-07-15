package definition

import "github.com/google/uuid"

type CreateService interface {
	Card(CreateCardRequest) CardResponse
	User(CreateUserRequest) UserResponse
}

type GetService interface {
	Card(CardRequest) CardResponse
	Cards(CardsRequest) CardsResponse
	User(UserRequest) UserResponse
	Users(UsersRequest) UsersResponse
	HomePage(UserRequest) HomePageResponse
}

type CreateCardRequest struct {
	Title   string
	Content string
	Creator string
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

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
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

type UsersRequest struct{}

type UsersResponse struct {
	Users []UserResponse
}

type CardsRequest struct{}

type CardsResponse struct {
	Cards []CardResponse
}

type HomePageResponse struct {
	User  *UserResponse
	Pairs []struct {
		User *UserResponse
		Card *CardResponse
	}
}
