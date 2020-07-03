package glance

import (
	"context"

	"github.com/tmitchel/glance/generated"
)

type CardService struct{}

func (CardService) Card(ctx context.Context, r generated.CardRequest) (*generated.CardResponse, error) {
	return nil, nil
}

func (CardService) Cards(ctx context.Context, r generated.CardsRequest) (*generated.CardsResponse, error) {
	return &generated.CardsResponse{
		Cards: []generated.CardResponse{},
	}, nil
}
