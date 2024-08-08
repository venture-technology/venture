package service

import (
	"context"

	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"
)

type ChildService struct {
	childrepository repository.IChildRepository
}

func NewChildService(childrepository repository.IChildRepository) *ChildService {
	return &ChildService{
		childrepository: childrepository,
	}
}

func (cs *ChildService) CreateChild(ctx context.Context, child *models.Child) error {
	return cs.childrepository.CreateChild(ctx, child)
}

func (cs *ChildService) GetChild(ctx context.Context, rg *string) (*models.Child, error) {
	return cs.childrepository.GetChild(ctx, rg)
}

func (cs *ChildService) FindAllChildren(ctx context.Context, cpf *string) ([]models.Child, error) {
	return cs.childrepository.FindAllChildren(ctx, cpf)
}

func (cs *ChildService) UpdateChild(ctx context.Context, child *models.Child) error {
	return cs.childrepository.UpdateChild(ctx, child)
}

func (cs *ChildService) DeleteChild(ctx context.Context, rg *string) error {
	return cs.childrepository.DeleteChild(ctx, rg)
}
