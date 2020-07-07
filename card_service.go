package glance

import (
	"context"
	"time"

	"github.com/tmitchel/glance/generated"
)

type CreateService struct{}

func (CreateService) Card(ctx context.Context, r generated.CreateCardRequest) (*generated.CardResponse, error) {
	return &generated.CardResponse{
		Title:     r.Title,
		Content:   r.Content,
		CreatedAt: time.Now().String(),
	}, nil
}

func (CreateService) User(ctx context.Context, r generated.CreateUserRequest) (*generated.UserResponse, error) {
	return &generated.UserResponse{
		Name:  r.Name,
		Email: r.Email,
	}, nil
}

type GetService struct{}

func (GetService) Card(ctx context.Context, r generated.CardRequest) (*generated.CardResponse, error) {
	return nil, nil
}

func (GetService) Cards(ctx context.Context, r generated.CardsRequest) (*generated.CardsResponse, error) {
	return &generated.CardsResponse{
		Cards: []generated.CardResponse{},
	}, nil
}

func (GetService) User(ctx context.Context, r generated.UserRequest) (*generated.UserResponse, error) {
	return nil, nil
}
