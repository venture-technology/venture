package repository

import "database/sql"

type IMapsRepository interface {
}

type MapsRepository struct {
	db *sql.DB
}

func NewMapsRepository(db *sql.DB) *MapsRepository {
	return &MapsRepository{
		db: db,
	}
}
