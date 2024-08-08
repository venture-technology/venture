package service

import (
	"context"

	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"
)

type InviteService struct {
	inviterepository repository.IInviteRepository
}

func NewInviteService(repo repository.IInviteRepository) *InviteService {
	return &InviteService{
		inviterepository: repo,
	}
}

func (i *InviteService) InviteDriver(ctx context.Context, invite *models.Invite) error {
	return i.inviterepository.InviteDriver(ctx, invite)
}

func (i *InviteService) ReadInvite(ctx context.Context, invite_id *int) (*models.Invite, error) {
	return i.inviterepository.ReadInvite(ctx, invite_id)
}

func (i *InviteService) FindAllInvitesDriverAccount(ctx context.Context, cnh *string) ([]models.Invite, error) {
	return i.inviterepository.FindAllInvitesDriverAccount(ctx, cnh)
}

func (i *InviteService) AcceptedInvite(ctx context.Context, invite *models.Invite) error {

	err := i.CreatePartner(ctx, invite)

	if err != nil {
		return err
	}

	return i.inviterepository.AcceptedInvite(ctx, &invite.ID)
}

func (i *InviteService) DeclineInvite(ctx context.Context, invite_id *int) error {
	return i.inviterepository.DeclineInvite(ctx, invite_id)
}

// Request in AccountManager to verify if school have the driver like employee. If they are partners, Employee is true, otherwise false.
func (i *InviteService) IsPartner(ctx context.Context, invite *models.Invite) (bool, error) {
	return i.inviterepository.IsPartner(ctx, invite)
}

// create partner between school and driver, then driver accepted invite, sending request to account manager
func (i *InviteService) CreatePartner(ctx context.Context, invite *models.Invite) error {
	return i.inviterepository.CreatePartner(ctx, invite)
}
