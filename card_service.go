package glance

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tmitchel/glance/generated"
)

// CreateService wraps the database and implements the CreateService
// interface from oto.
type CreateService struct {
	db Database
}

// NewCreateService returns a CreateService wrapping the provided database.
func NewCreateService(db Database) (*CreateService, error) {
	return &CreateService{db}, nil
}

// Card creates and saves a new card.
func (c *CreateService) Card(ctx context.Context, r generated.CreateCardRequest) (*generated.CardResponse, error) {
	card := Card{
		ID:        uuid.New(),
		Title:     r.Title,
		Content:   r.Content,
		Creator:   uuid.MustParse(r.Creator),
		CreatedAt: time.Now(),
	}
	err := c.db.CreateCard(card)
	if err != nil {
		return &generated.CardResponse{Error: err.Error()}, err
	}

	return cardResponse(card), nil
}

// User creates and saves a new user.
func (c *CreateService) User(ctx context.Context, r generated.CreateUserRequest) (*generated.UserResponse, error) {
	user := User{
		ID:       uuid.New(),
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
	err := c.db.CreateUser(user)
	if err != nil {
		return &generated.UserResponse{Error: err.Error()}, err
	}

	return userResponse(user), nil
}

func (c *CreateService) ClaimCard(ctx context.Context, r generated.ClaimRequest) (*generated.CardResponse, error) {
	card, err := c.db.ClaimCard(r.UserID, r.CardID)
	if err != nil {
		return &generated.CardResponse{Error: err.Error()}, err
	}

	return cardResponse(card), nil
}

func (c *CreateService) UpdateStatus(ct context.Context, r generated.NewStatusRequest) (*generated.CardResponse, error) {
	card, err := c.db.UpdateStatus(r.Card, r.User, r.Status)
	if err != nil {
		return &generated.CardResponse{Error: err.Error()}, err
	}

	return cardResponse(card), nil
}

// GetService wraps the database and implements the GetService
// interface from oto.
type GetService struct {
	db Database
}

// NewGetService returns a GetService wrapping the provided database.
func NewGetService(db Database) (*GetService, error) {
	return &GetService{db}, nil
}

// Card returns a single card by ID.
func (g *GetService) Card(ctx context.Context, r generated.CardRequest) (*generated.CardResponse, error) {
	card, err := g.db.GetCard(r.ID)
	if err != nil {
		return &generated.CardResponse{Error: err.Error()}, err
	}

	return cardResponse(card), nil
}

// Cards returns all cards.
func (g *GetService) Cards(ctx context.Context, r generated.EmptyRequest) (*generated.CardsResponse, error) {
	cards, err := g.db.GetUnclaimedCards()
	if err != nil {
		return &generated.CardsResponse{Error: err.Error()}, err
	}

	return cardsResponse(cards), nil
}

// User returns a user by email.
func (g *GetService) User(ctx context.Context, r generated.UserRequest) (*generated.UserResponse, error) {
	user, err := g.db.GetUser(r.Email)
	if err != nil {
		return &generated.UserResponse{Error: err.Error()}, err
	}

	return userResponse(user), nil
}

// Users returns all Users.
func (g *GetService) Users(ctx context.Context, r generated.EmptyRequest) (*generated.UsersResponse, error) {
	users, err := g.db.GetUsers()
	if err != nil {
		return &generated.UsersResponse{Error: err.Error()}, err
	}

	return usersResponse(users), nil
}

func (g *GetService) HomePage(ctx context.Context, r generated.UserRequest) (*generated.HomePageResponse, error) {
	user, err := g.db.GetUser(r.Email)
	if err != nil {
		return &generated.HomePageResponse{Error: err.Error()}, err
	}

	users, err := g.db.GetUsers()
	if err != nil {
		return &generated.HomePageResponse{Error: err.Error()}, err
	}

	pairs := make([]struct {
		User *generated.UserResponse
		Card *generated.CardResponse
	}, len(users))

	for i, u := range users {
		card, _ := g.db.GetUserCurrentCard(u.ID.String())
		pairs[i] = struct {
			User *generated.UserResponse
			Card *generated.CardResponse
		}{User: userResponse(u), Card: cardResponse(card)}
	}

	return &generated.HomePageResponse{
		User:  userResponse(user),
		Pairs: pairs,
	}, nil
}

func cardResponse(c Card) *generated.CardResponse {
	return &generated.CardResponse{
		ID:        c.ID,
		Title:     c.Title,
		Content:   c.Content,
		Creator:   c.Creator,
		Claimed:   c.Claimed,
		CreatedAt: c.Creator.String(),
		Status:    c.Status,
	}
}

func cardsResponse(c []Card) *generated.CardsResponse {
	cards := make([]generated.CardResponse, len(c))
	for i, card := range c {
		cards[i] = *cardResponse(card)
	}
	return &generated.CardsResponse{Cards: cards}
}

func userResponse(u User) *generated.UserResponse {
	return &generated.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func usersResponse(u []User) *generated.UsersResponse {
	users := make([]generated.UserResponse, len(u))
	for i, user := range u {
		users[i] = *userResponse(user)
	}
	return &generated.UsersResponse{Users: users}
}
