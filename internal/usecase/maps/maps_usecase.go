package maps

import (
	"context"

	"github.com/venture-technology/venture/internal/domain/adapter"
	"github.com/venture-technology/venture/internal/usecase"
)

type MapsUseCase struct {
	googleAdapter adapter.IGoogleAdapter
}

func NewMapsUseCase(googleAdapter adapter.IGoogleAdapter) *MapsUseCase {
	return &MapsUseCase{
		googleAdapter: googleAdapter,
	}
}

func (mu *MapsUseCase) CalculatePrice(ctx context.Context, origin, destination string, amount float64) (*float64, error) {
	km, err := mu.googleAdapter.GetDistance(origin, destination)

	if err != nil {
		return nil, err
	}

	value := usecase.CalculateContract(*km, amount)

	return &value, nil
}
