package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type ListChildrenUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListChildrenUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListChildrenUseCase {
	return &ListChildrenUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lcuc *ListChildrenUseCase) ListChildren(cpf *string) ([]value.ListChild, error) {
	children, err := lcuc.repositories.ChildRepository.FindAll(cpf)
	if err != nil {
		return []value.ListChild{}, err
	}
	response := []value.ListChild{}
	for _, child := range children {
		response = append(response, buildListChild(&child))
	}
	return response, nil
}

func buildListChild(child *entity.Child) value.ListChild {
	return value.ListChild{
		ID:           child.ID,
		Name:         child.Name,
		Period:       child.Shift,
		ProfileImage: child.ProfileImage,
	}
}
