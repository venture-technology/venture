package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
)

type IInviteRepository interface {
	Create(ctx context.Context, invite *entity.Invite) error
	Get(ctx context.Context, id uuid.UUID) (*entity.Invite, error)
	FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Invite, error)
	FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Invite, error)
	Accept(ctx context.Context, id uuid.UUID) error
	Decline(ctx context.Context, id uuid.UUID) error
}

type InviteRepository struct {
	db *sql.DB
}

func NewInviteRepository(db *sql.DB) *InviteRepository {
	return &InviteRepository{
		db: db,
	}
}

func (ir *InviteRepository) Create(ctx context.Context, invite *entity.Invite) error {
	sqlQuery := `INSERT INTO invites (id, requester, guester, status) VALUES ($1, $2, $3, $4)`
	_, err := ir.db.Exec(sqlQuery, invite.ID, invite.School.CNPJ, invite.Driver.CNH, "pending")
	return err
}

func (i *InviteRepository) Get(ctx context.Context, inviteID uuid.UUID) (*entity.Invite, error) {
	sqlQuery := `
        SELECT 
            i.id, 
            s.cnpj, 
            s.name AS school_name, 
            s.email AS school_email, 
            d.cnh, 
            d.name AS driver_name, 
            d.email AS driver_email, 
            i.status 
        FROM 
            invites i
        JOIN 
            schools s ON i.requester = s.cnpj
        JOIN 
            drivers d ON i.guester = d.cnh
        WHERE 
            i.id = $1 
        LIMIT 1;
    `

	var invite entity.Invite
	err := i.db.QueryRowContext(ctx, sqlQuery, inviteID).Scan(
		&invite.ID,
		&invite.School.CNPJ,
		&invite.School.Name,
		&invite.School.Email,
		&invite.Driver.CNH,
		&invite.Driver.Name,
		&invite.Driver.Email,
		&invite.Status,
	)
	if err != nil {
		return nil, err
	}

	return &invite, nil
}

func (i *InviteRepository) FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Invite, error) {
	sqlQuery := `
        SELECT 
            i.id, 
            s.cnpj, 
            s.name AS school_name, 
            s.email AS school_email, 
            d.cnh, 
            d.name AS driver_name, 
            d.email AS driver_email, 
            i.status
        FROM 
            invites i
        JOIN 
            schools s ON i.requester = s.cnpj
        JOIN 
            drivers d ON i.guester = d.cnh
        WHERE 
            i.status = 'pending' 
            AND i.guester = $1;
    `

	rows, err := i.db.QueryContext(ctx, sqlQuery, *cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invites []entity.Invite

	for rows.Next() {
		var invite entity.Invite
		err := rows.Scan(
			&invite.ID,
			&invite.School.CNPJ,
			&invite.School.Name,
			&invite.School.Email,
			&invite.Driver.CNH,
			&invite.Driver.Name,
			&invite.Driver.Email,
			&invite.Status,
		)
		if err != nil {
			return nil, err
		}
		invites = append(invites, invite)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return invites, nil
}

func (i *InviteRepository) FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Invite, error) {
	sqlQuery := `
        SELECT 
            i.id, 
            s.cnpj, 
            s.name AS school_name, 
            s.email AS school_email, 
            d.cnh, 
            d.name AS driver_name, 
            d.email AS driver_email, 
            i.status
        FROM 
            invites i
        JOIN 
            schools s ON i.requester = s.cnpj
        JOIN 
            drivers d ON i.guester = d.cnh
        WHERE 
            i.status = 'pending' 
            AND i.requester = $1;
    `

	rows, err := i.db.QueryContext(ctx, sqlQuery, *cnpj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invites []entity.Invite

	for rows.Next() {
		var invite entity.Invite
		err := rows.Scan(
			&invite.ID,
			&invite.School.CNPJ,
			&invite.School.Name,
			&invite.School.Email,
			&invite.Driver.CNH,
			&invite.Driver.Name,
			&invite.Driver.Email,
			&invite.Status,
		)
		if err != nil {
			return nil, err
		}
		invites = append(invites, invite)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return invites, nil
}

func (i *InviteRepository) Accept(ctx context.Context, id uuid.UUID) error {
	sqlQuery := `
        WITH updated_invite AS (
            UPDATE invites
            SET status = 'accepted'
            WHERE id = $1
            RETURNING requester, guester
        )
        INSERT INTO partners (driver_id, school_id)
        SELECT guester, requester
        FROM updated_invite;
    `

	_, err := i.db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}

func (ir *InviteRepository) Decline(ctx context.Context, id uuid.UUID) error {
	tx, err := ir.db.Begin()
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
	_, err = tx.Exec("DELETE FROM invites WHERE id = $1", id)
	return err
}
