package service

import (
	"context"

	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"
)

type InviteService struct {
	inviterepository   repository.IInviteRepository
	partnersrepository repository.IPartnerRepository
}

func NewInviteService(inviterepository repository.IInviteRepository, partnersrepository repository.IPartnerRepository) *InviteService {
	return &InviteService{
		inviterepository:   inviterepository,
		partnersrepository: partnersrepository,
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

	partner := models.Partner{
		Driver: invite.Driver,
		School: invite.School,
	}

	err := i.partnersrepository.CreatePartners(ctx, &partner)

	if err != nil {
		return err
	}

	return i.inviterepository.AcceptedInvite(ctx, &invite.ID)
}

func (i *InviteService) DeclineInvite(ctx context.Context, invite_id *int) error {
	return i.inviterepository.DeclineInvite(ctx, invite_id)
}
