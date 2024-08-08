package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/venture-technology/venture/models"
	"github.com/venture-technology/venture/utils"
)

type SchoolRepositoryInterface interface {
	CreateSchool(ctx context.Context, school *models.School) error
	ReadSchool(ctx context.Context, cnpj *string) (*models.School, error)
	ReadAllSchools(ctx context.Context) ([]models.School, error)
	UpdateSchool(ctx context.Context, school *models.School) error
	DeleteSchool(ctx context.Context, cnpj *string) error
	AuthSchool(ctx context.Context, school *models.School) (*models.School, error)
}

type SchoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{
		db: db,
	}
}

func (s *SchoolRepository) CreateSchool(ctx context.Context, school *models.School) error {
	sqlQuery := `INSERT INTO schools (name, cnpj, email, password, street, number, zip) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(sqlQuery, school.Name, school.CNPJ, school.Email, school.Password, school.Street, school.Number, school.ZIP)
	return err
}

func (s *SchoolRepository) ReadSchool(ctx context.Context, cnpj *string) (*models.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip FROM schools WHERE cnpj = $1 LIMIT 1`
	var school models.School
	err := s.db.QueryRow(sqlQuery, *cnpj).Scan(
		&school.ID,
		&school.Name,
		&school.CNPJ,
		&school.Email,
		&school.Street,
		&school.Number,
		&school.ZIP,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &school, nil
}

func (s *SchoolRepository) ReadAllSchools(ctx context.Context) ([]models.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip FROM schools`

	rows, err := s.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []models.School

	for rows.Next() {
		var school models.School
		err := rows.Scan(&school.ID, &school.Name, &school.CNPJ, &school.Email, &school.Street, &school.Number, &school.ZIP)
		if err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schools, nil
}

func (s *SchoolRepository) UpdateSchool(ctx context.Context, school *models.School) error {
	sqlQuery := `SELECT name, email, password, street, number, zip FROM schools WHERE cnpj = $1 LIMIT 1`

	var currentSchool models.School
	err := s.db.QueryRow(sqlQuery, &school.CNPJ).Scan(
		&currentSchool.Name,
		&currentSchool.Email,
		&currentSchool.Password,
		&currentSchool.Street,
		&currentSchool.Number,
		&currentSchool.ZIP,
	)
	if err != nil || err == sql.ErrNoRows {
		return err
	}

	if school.Name != "" && school.Name != currentSchool.Name {
		currentSchool.Name = school.Name
	}
	if school.Email != "" && school.Email != currentSchool.Email {
		currentSchool.Email = school.Email
	}
	if school.Password != "" && school.Password != currentSchool.Password {
		school.Password = utils.HashPassword(school.Password)
		currentSchool.Password = school.Password
	}
	if school.Street != "" && school.Street != currentSchool.Street {
		currentSchool.Street = school.Street
	}
	if school.Number != "" && school.Number != currentSchool.Number {
		currentSchool.Number = school.Number
	}
	if school.ZIP != "" && school.ZIP != currentSchool.ZIP {
		currentSchool.ZIP = school.ZIP
	}

	sqlQueryUpdate := `UPDATE schools SET name = $1, email = $2, password = $3, street = $4, number = $5, zip = $6 WHERE cnpj = $7`
	_, err = s.db.ExecContext(ctx, sqlQueryUpdate, currentSchool.Name, currentSchool.Email, currentSchool.Password, currentSchool.Street, currentSchool.Number, currentSchool.ZIP, &school.CNPJ)
	return err

}

func (s *SchoolRepository) DeleteSchool(ctx context.Context, cnpj *string) error {
	tx, err := s.db.Begin()
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
	_, err = tx.Exec("DELETE FROM schools WHERE cnpj = $1", cnpj)
	return err
}

func (s *SchoolRepository) AuthSchool(ctx context.Context, school *models.School) (*models.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, password FROM schools WHERE email = $1 LIMIT 1`
	var schoolData models.School
	err := s.db.QueryRow(sqlQuery, school.Email).Scan(
		&schoolData.ID,
		&schoolData.Name,
		&schoolData.CNPJ,
		&schoolData.Email,
		&schoolData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := schoolData.Password == school.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	schoolData.Password = ""
	return &schoolData, nil
}
