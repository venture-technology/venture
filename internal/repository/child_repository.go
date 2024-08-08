package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/models"
)

type IChildRepository interface {
	CreateChild(ctx context.Context, child *models.Child) error
	GetChild(ctx context.Context, rg *string) (*models.Child, error)
	FindAllChildren(ctx context.Context, cpf *string) ([]models.Child, error)
	UpdateChild(ctx context.Context, child *models.Child) error
	DeleteChild(ctx context.Context, rg *string) error
}

type ChildRepository struct {
	db *sql.DB
}

func NewChildRepository(conn *sql.DB) *ChildRepository {
	return &ChildRepository{
		db: conn,
	}
}

func (cr *ChildRepository) CreateChild(ctx context.Context, child *models.Child) error {
	sqlQuery := `INSERT INTO children (name, rg, responsible_id,shift) VALUES ($1, $2, $3, $4)`
	_, err := cr.db.Exec(sqlQuery, child.Name, child.RG, child.Responsible.CPF, child.Shift)
	return err
}

func (cr *ChildRepository) GetChild(ctx context.Context, rg *string) (*models.Child, error) {
	sqlQuery := `SELECT id, name, rg, responsible_id, shift FROM children WHERE rg = $1 LIMIT 1`
	var child models.Child
	err := cr.db.QueryRow(sqlQuery, *rg).Scan(&child.ID, &child.Name, &child.RG, &child.Responsible.CPF, &child.Shift)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &child, nil
}

func (cr *ChildRepository) FindAllChildren(ctx context.Context, cpf *string) ([]models.Child, error) {
	sqlQuery := `SELECT id, name, rg, responsible_id, shift FROM children WHERE responsible_id = $1`
	rows, err := cr.db.Query(sqlQuery, cpf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var children []models.Child
	for rows.Next() {
		var child models.Child
		err := rows.Scan(
			&child.ID,
			&child.Name,
			&child.RG,
			&child.Responsible.CPF,
			&child.Shift,
		)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return children, nil
}

func (cr *ChildRepository) UpdateChild(ctx context.Context, child *models.Child) error {
	sqlQuery := `SELECT name, shift FROM children WHERE rg = $1 LIMIT 1`
	var currentChild models.Child
	err := cr.db.QueryRow(sqlQuery, child.RG).Scan(&currentChild.Name)
	if err != nil || err == sql.ErrNoRows {
		return err
	}
	if child.Name != "" && child.Name != currentChild.Name {
		currentChild.Name = child.Name
	}
	if child.Shift != "" && child.Shift != currentChild.Shift {
		currentChild.Shift = child.Shift
	}
	sqlQueryUpdate := `UPDATE children SET name = $1, shift = $2 WHERE rg = $3`
	_, err = cr.db.ExecContext(ctx, sqlQueryUpdate, currentChild.Name, currentChild.Shift, child.RG)
	return err
}

func (cr *ChildRepository) DeleteChild(ctx context.Context, rg *string) error {
	tx, err := cr.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM children WHERE rg = $1", *rg)
	return err
}
