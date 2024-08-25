package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/internal/entity"
)

type ISchoolRepository interface {
	Create(ctx context.Context, school *entity.School) error
	Get(ctx context.Context, cnpj *string) (*entity.School, error)
	FindAll(ctx context.Context) ([]entity.School, error)
	Update(ctx context.Context, currentSchool, school *entity.School) error
	Delete(ctx context.Context, cnpj *string) error
}

type SchoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{
		db: db,
	}
}

func (sr *SchoolRepository) Create(ctx context.Context, school *entity.School) error {
	return nil
}

func (sr *SchoolRepository) Get(ctx context.Context, cnpj *string) (*entity.School, error) {
	return nil, nil
}

func (sr *SchoolRepository) FindAll(ctx context.Context) ([]entity.School, error) {
	return nil, nil
}

func (sr *SchoolRepository) Update(ctx context.Context, currentSchool, school *entity.School) error {
	return nil
}

func (sr *SchoolRepository) Delete(ctx context.Context, cnpj *string) error {
	return nil
}
