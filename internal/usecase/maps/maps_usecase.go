package maps

import (
	"context"

	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase"
)

type MapsUseCase struct {
	mapsRepository repository.IMapsRepository
}

func NeWMapsUseCase(mapsRepository repository.IMapsRepository) *MapsUseCase {
	return &MapsUseCase{
		mapsRepository: mapsRepository,
	}
}

func (mu *MapsUseCase) CalculatePrice(ctx context.Context, origin, destination string, amount float64) (*float64, error) {

	km, err := usecase.GetDistance(origin, destination)

	if err != nil {
		return nil, err
	}

	value := usecase.CalculateContract(*km, amount)

	return &value, nil

}
