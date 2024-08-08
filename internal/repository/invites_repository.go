package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/models"
)

type IInviteRepository interface {
	InviteDriver(ctx context.Context, invite *models.Invite) error
	ReadInvite(ctx context.Context, invite_id *int) (*models.Invite, error)
	FindAllInvitesDriverAccount(ctx context.Context, cnh *string) ([]models.Invite, error)
	AcceptedInvite(ctx context.Context, invite_id *int) error
	DeclineInvite(ctx context.Context, invite_id *int) error
}

type InviteRepository struct {
	db *sql.DB
}

func NewInviteRepository(db *sql.DB) *InviteRepository {
	return &InviteRepository{
		db: db,
	}
}

func (i *InviteRepository) InviteDriver(ctx context.Context, invite *models.Invite) error {
	sqlQuery := `INSERT INTO invites (requester, school, email_school, guest, driver, email_driver, status) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := i.db.Exec(sqlQuery, invite.School.CNPJ, invite.School.Name, invite.School.Email, invite.Driver.CNH, invite.Driver.Name, invite.Driver.Email, "pending")
	return err
}

func (i *InviteRepository) ReadInvite(ctx context.Context, invite_id *int) (*models.Invite, error) {
	sqlQuery := `SELECT invite_id, requester, school, email_school, guest, driver, email_driver, status FROM invites WHERE invite_id = $1 LIMIT 1`
	var invite models.Invite
	err := i.db.QueryRow(sqlQuery, *invite_id).Scan(
		&invite.ID,
		&invite.School.CNPJ,
		&invite.School.Name,
		&invite.School.Email,
		&invite.Driver.CNH,
		&invite.Driver.Name,
		&invite.Driver.Email,
		&invite.Status,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &invite, nil
}

func (i *InviteRepository) FindAllInvitesDriverAccount(ctx context.Context, cnh *string) ([]models.Invite, error) {
	sqlQuery := `SELECT invite_id, school, requester, email_school, driver, guest, email_driver, status FROM invites WHERE status = 'pending' AND guest = $1`

	rows, err := i.db.Query(sqlQuery, *cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invites []models.Invite

	for rows.Next() {
		var invite models.Invite
		err := rows.Scan(&invite.ID, &invite.School.Name, &invite.School.CNPJ, &invite.School.Email, &invite.Driver.Name, &invite.Driver.CNH, &invite.Driver.Email, &invite.Status)
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

func (i *InviteRepository) AcceptedInvite(ctx context.Context, invite_id *int) error {
	sqlQuery := `UPDATE invites SET status = 'accepted' WHERE invite_id = $1`
	_, err := i.db.Exec(sqlQuery, invite_id)
	return err
}

func (i *InviteRepository) DeclineInvite(ctx context.Context, invite_id *int) error {
	tx, err := i.db.Begin()
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
	_, err = tx.Exec("DELETE FROM invites WHERE invite_id = $1", invite_id)
	return err

}
