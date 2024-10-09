package invite

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
	"go.uber.org/zap"
)

type InviteUseCase struct {
	inviteRepository  repository.IInviteRepository
	partnerRepository repository.IPartnerRepository
	logger            *zap.Logger
}

func NewInviteUseCase(ir repository.IInviteRepository, pr repository.IPartnerRepository, logger *zap.Logger) *InviteUseCase {
	return &InviteUseCase{
		inviteRepository:  ir,
		partnerRepository: pr,
		logger:            logger,
	}
}

func (iu *InviteUseCase) Create(ctx context.Context, invite *entity.Invite) error {

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	invite.ID = id

	log.Print("validating partner")
	isPartner, err := iu.partnerRepository.IsPartner(ctx, &invite.Driver.CNH, &invite.School.CNPJ)

	if err != nil {
		return err
	}

	if isPartner {
		return fmt.Errorf("it wasnt created invite because they are partners")
	}

	log.Print("creating invite")
	return iu.inviteRepository.Create(ctx, invite)

}

func (iu *InviteUseCase) Get(ctx context.Context, id uuid.UUID) (*entity.Invite, error) {
	return iu.inviteRepository.Get(ctx, id)
}

func (iu *InviteUseCase) FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Invite, error) {
	return iu.inviteRepository.FindAllByCnh(ctx, cnh)
}

func (iu *InviteUseCase) FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Invite, error) {
	return iu.inviteRepository.FindAllByCnpj(ctx, cnpj)
}

func (iu *InviteUseCase) Accept(ctx context.Context, id uuid.UUID) error {
	return iu.inviteRepository.Accept(ctx, id)
}

func (iu *InviteUseCase) Decline(ctx context.Context, id uuid.UUID) error {
	return iu.inviteRepository.Decline(ctx, id)
}
