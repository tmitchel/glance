package glance

import (
	"context"
	"time"

	"github.com/tmitchel/glance/generated"
)

type CardService struct{}

func (CardService) New(ctx context.Context, r generated.CreateRequest) (*generated.CardResponse, error) {
	return &generated.CardResponse{
		Title:     r.Title,
		Content:   r.Content,
		CreatedAt: time.Now().String(),
	}, nil
}

func (CardService) Card(ctx context.Context, r generated.CardRequest) (*generated.CardResponse, error) {
	return nil, nil
}

func (CardService) Cards(ctx context.Context, r generated.CardsRequest) (*generated.CardsResponse, error) {
	return &generated.CardsResponse{
		Cards: []generated.CardResponse{},
	}, nil
}
