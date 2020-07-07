package glance

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tmitchel/glance/generated"
)

type CreateService struct {
	db Database
}

func NewCreateService(db Database) (*CreateService, error) {
	return &CreateService{db}, nil
}

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

type GetService struct {
	db Database
}

func NewGetService(db Database) (*GetService, error) {
	return &GetService{db}, nil
}

func (g *GetService) Card(ctx context.Context, r generated.CardRequest) (*generated.CardResponse, error) {
	card, err := g.db.GetCard(r.ID)
	if err != nil {
		return &generated.CardResponse{Error: err.Error()}, err
	}

	return cardResponse(card), nil
}

func (g *GetService) Cards(ctx context.Context, r generated.CardsRequest) (*generated.CardsResponse, error) {
	cards, err := g.db.GetCards()
	if err != nil {
		return &generated.CardsResponse{Error: err.Error()}, err
	}

	return cardsResponse(cards), nil
}

func (g *GetService) User(ctx context.Context, r generated.UserRequest) (*generated.UserResponse, error) {
	user, err := g.db.GetUser(r.Email)
	if err != nil {
		return &generated.UserResponse{Error: err.Error()}, err
	}

	return userResponse(user), nil
}

func cardResponse(c Card) *generated.CardResponse {
	return &generated.CardResponse{
		ID:        c.ID,
		Title:     c.Title,
		Content:   c.Content,
		Creator:   c.Creator,
		Volunteer: c.Volunteer,
		CreatedAt: c.Creator.String(),
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
