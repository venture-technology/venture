package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/internal/entity"
	"go.uber.org/zap"
)

type IChildRepository interface {
	Create(ctx context.Context, child *entity.Child) error
	Get(ctx context.Context, rg *string) (*entity.Child, error)
	FindAll(ctx context.Context, cpf *string) ([]entity.Child, error)
	Update(ctx context.Context, child *entity.Child) error
	Delete(ctx context.Context, rg *string) error
}

type ChildRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewChildRepository(conn *sql.DB, logger *zap.Logger) *ChildRepository {
	return &ChildRepository{
		db:     conn,
		logger: logger,
	}
}

func (cr *ChildRepository) Create(ctx context.Context, child *entity.Child) error {
	sqlQuery := `INSERT INTO children (name, rg, responsible_id, shift, profile_image) VALUES ($1, $2, $3, $4. $5)`
	_, err := cr.db.Exec(sqlQuery, child.Name, child.RG, child.Responsible.CPF, child.Shift, child.ProfileImage)
	return err
}

func (cr *ChildRepository) Get(ctx context.Context, rg *string) (*entity.Child, error) {
	sqlQuery := `SELECT id, name, rg, responsible_id, shift, profile_image FROM children WHERE rg = $1 LIMIT 1`
	var child entity.Child
	err := cr.db.QueryRow(sqlQuery, *rg).Scan(&child.ID, &child.Name, &child.RG, &child.Responsible.CPF, &child.Shift, &child.ProfileImage)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &child, nil
}

func (cr *ChildRepository) FindAll(ctx context.Context, cpf *string) ([]entity.Child, error) {
	sqlQuery := `SELECT id, name, rg, responsible_id, shift, profile_image FROM children WHERE responsible_id = $1`
	rows, err := cr.db.Query(sqlQuery, cpf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var children []entity.Child
	for rows.Next() {
		var child entity.Child
		err := rows.Scan(
			&child.ID,
			&child.Name,
			&child.RG,
			&child.Responsible.CPF,
			&child.Shift,
			&child.ProfileImage,
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

func (cr *ChildRepository) Update(ctx context.Context, child *entity.Child) error {
	sqlQuery := `SELECT name, shift, profile_image FROM children WHERE rg = $1 LIMIT 1`
	var currentChild entity.Child
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
	if child.ProfileImage != "" && child.ProfileImage != currentChild.ProfileImage {
		currentChild.ProfileImage = child.ProfileImage
	}
	sqlQueryUpdate := `UPDATE children SET name = $1, shift = $2, profile_image = $3 WHERE rg = $4`
	_, err = cr.db.ExecContext(ctx, sqlQueryUpdate, currentChild.Name, currentChild.Shift, currentChild.ProfileImage, child.RG)
	return err
}

func (cr *ChildRepository) Delete(ctx context.Context, rg *string) error {
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
