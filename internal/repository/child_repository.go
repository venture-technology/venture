package repository

import "database/sql"

type IChildRepository interface {
	Create() error
	Get() error
	Update() error
	Delete() error
}

type ChildRepository struct {
	db *sql.DB
}

func NewChildRepository(db *sql.DB) *ChildRepository {
	return &ChildRepository{
		db: db,
	}
}

func (cr *ChildRepository) Create() error {
	return nil
}

func (cr *ChildRepository) Get() error {
	return nil
}

func (cr *ChildRepository) Update() error {
	return nil
}

func (cr *ChildRepository) Delete() error {
	return nil
}
