package child

import (
	"context"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
	"go.uber.org/zap"
)

type ChildUseCase struct {
	childRepository repository.IChildRepository
	logger          *zap.Logger
}

func NewChildUseCase(childRepository repository.IChildRepository, logger *zap.Logger) *ChildUseCase {
	return &ChildUseCase{
		childRepository: childRepository,
		logger:          logger,
	}
}

func (cu *ChildUseCase) Create(ctx context.Context, child *entity.Child) error {
	return cu.childRepository.Create(ctx, child)
}

func (cu *ChildUseCase) Get(ctx context.Context, rg *string) (*entity.Child, error) {
	return cu.childRepository.Get(ctx, rg)
}

func (cu *ChildUseCase) FindAll(ctx context.Context, cpf *string) ([]entity.Child, error) {
	return cu.childRepository.FindAll(ctx, cpf)
}

func (cu *ChildUseCase) Update(ctx context.Context, child *entity.Child) error {
	return cu.childRepository.Update(ctx, child)
}

func (cu *ChildUseCase) Delete(ctx context.Context, rg *string) error {
	return cu.childRepository.Delete(ctx, rg)
}
