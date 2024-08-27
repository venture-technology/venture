package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/pkg/utils"
)

type ISchoolRepository interface {
	Create(ctx context.Context, school *entity.School) error
	Get(ctx context.Context, cnpj *string) (*entity.School, error)
	FindAll(ctx context.Context) ([]entity.School, error)
	Update(ctx context.Context, school *entity.School) error
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
	sqlQuery := `INSERT INTO schools (name, cnpj, email, password, street, number, zip, phone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := sr.db.Exec(sqlQuery, school.Name, school.CNPJ, school.Email, school.Password, school.Address.Street, school.Address.Number, school.Address.ZIP, school.Phone)
	return err
}

func (sr *SchoolRepository) Get(ctx context.Context, cnpj *string) (*entity.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip, phone FROM schools WHERE cnpj = $1 LIMIT 1`
	var school entity.School
	err := sr.db.QueryRow(sqlQuery, *cnpj).Scan(
		&school.ID,
		&school.Name,
		&school.CNPJ,
		&school.Email,
		&school.Address.Street,
		&school.Address.Number,
		&school.Address.ZIP,
		&school.Phone,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &school, nil
}

func (sr *SchoolRepository) FindAll(ctx context.Context) ([]entity.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip, phone FROM schools`

	rows, err := sr.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []entity.School

	for rows.Next() {
		var school entity.School
		err := rows.Scan(&school.ID, &school.Name, &school.CNPJ, &school.Email, &school.Address.Street, &school.Address.Number, &school.Address.ZIP, &school.Phone)
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

func (sr *SchoolRepository) Update(ctx context.Context, school *entity.School) error {
	sqlQuery := `SELECT name, email, password, street, number, zip, phone FROM schools WHERE cnpj = $1 LIMIT 1`

	var currentSchool entity.School
	err := sr.db.QueryRow(sqlQuery, &school.CNPJ).Scan(
		&currentSchool.Name,
		&currentSchool.Email,
		&currentSchool.Password,
		&currentSchool.Address.Street,
		&currentSchool.Address.Number,
		&currentSchool.Address.ZIP,
		&currentSchool.Phone,
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
	if school.Address.Street != "" && school.Address.Street != currentSchool.Address.Street {
		currentSchool.Address.Street = school.Address.Street
	}
	if school.Address.Number != "" && school.Address.Number != currentSchool.Address.Number {
		currentSchool.Address.Number = school.Address.Number
	}
	if school.Address.ZIP != "" && school.Address.ZIP != currentSchool.Address.ZIP {
		currentSchool.Address.ZIP = school.Address.ZIP
	}
	if school.Phone != "" && school.Phone != currentSchool.Phone {
		currentSchool.Phone = school.Phone
	}

	sqlQueryUpdate := `UPDATE schools SET name = $1, email = $2, password = $3, street = $4, number = $5, zip = $6, phone = $7 WHERE cnpj = $8`
	_, err = sr.db.ExecContext(ctx, sqlQueryUpdate, currentSchool.Name, currentSchool.Email, currentSchool.Password, currentSchool.Address.Street, currentSchool.Address.Number, currentSchool.Address.ZIP, currentSchool.Phone, &school.CNPJ)
	return err
}

func (sr *SchoolRepository) Delete(ctx context.Context, cnpj *string) error {
	tx, err := sr.db.Begin()
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
