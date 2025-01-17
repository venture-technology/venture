package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type ListDriverInvitesUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListDriverInvitesUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListDriverInvitesUseCase {
	return &ListDriverInvitesUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lsiuc *ListDriverInvitesUseCase) ListDriverInvites(cnh string) ([]value.DriverListInvite, error) {
	invites, err := lsiuc.repositories.InviteRepository.FindAllByCnh(cnh)
	if err != nil {
		return nil, err
	}
	response := []value.DriverListInvite{}
	for _, invite := range invites {
		response = append(response, buildDriverListInvites(invite))
	}
	return response, nil
}

func buildDriverListInvites(invites entity.Invite) value.DriverListInvite {
	return value.DriverListInvite{
		ID:           invites.ID,
		Email:        invites.School.Email,
		Name:         invites.School.Name,
		Phone:        invites.School.Phone,
		ProfileImage: invites.School.ProfileImage,
		Address: utils.BuildAddress(
			invites.School.Address.Street,
			invites.School.Address.Number,
			invites.School.Address.Complement,
			invites.School.Address.Zip,
		),
	}
}
