package school

import (
	"context"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
)

type SchoolUseCase struct {
	schoolRepository repository.ISchoolRepository
}

func NewSchoolUseCase(schoolRepository repository.ISchoolRepository) *SchoolUseCase {
	return &SchoolUseCase{
		schoolRepository: schoolRepository,
	}
}

func (su *SchoolUseCase) Create(ctx context.Context, school *entity.School) error {
	return su.schoolRepository.Create(ctx, school)
}

func (su *SchoolUseCase) Get(ctx context.Context, cnpj *string) (*entity.School, error) {
	return su.schoolRepository.Get(ctx, cnpj)
}

func (su *SchoolUseCase) FindAll(ctx context.Context) ([]entity.School, error) {
	return su.schoolRepository.FindAll(ctx)
}

func (su *SchoolUseCase) Update(ctx context.Context, school *entity.School) error {
	return su.schoolRepository.Update(ctx, school)
}

func (su *SchoolUseCase) Delete(ctx context.Context, cnpj *string) error {
	return su.schoolRepository.Delete(ctx, cnpj)
}