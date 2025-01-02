package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type ListSchoolInvitesUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListSchoolInvitesUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListSchoolInvitesUseCase {
	return &ListSchoolInvitesUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lsiuc *ListSchoolInvitesUseCase) ListSchoolInvites(cnpj string) ([]value.SchoolListInvite, error) {
	invites, err := lsiuc.repositories.InviteRepository.FindAllByCnpj(cnpj)
	if err != nil {
		return nil, err
	}
	response := []value.SchoolListInvite{}
	for _, invite := range invites {
		response = append(response, buildSchoolListInvites(invite))
	}
	return response, nil
}

func buildSchoolListInvites(invites entity.Invite) value.SchoolListInvite {
	return value.SchoolListInvite{
		ID:           invites.ID,
		Email:        invites.Driver.Email,
		Name:         invites.Driver.Name,
		Phone:        invites.Driver.Phone,
		ProfileImage: invites.Driver.ProfileImage,
	}
}
