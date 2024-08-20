package child

import "github.com/venture-technology/venture/internal/repository"

type ChildUseCase struct {
	childRepository repository.IChildRepository
}

func NewChildUseCase(childRepository repository.IChildRepository) *ChildUseCase {
	return &ChildUseCase{
		childRepository: childRepository,
	}
}
