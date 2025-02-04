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
	drivers, err := lsiuc.repositories.InviteRepository.FindAllByCnpj(cnpj)
	if err != nil {
		return nil, err
	}
	response := []value.SchoolListInvite{}
	for _, driver := range drivers {
		response = append(response, buildSchoolListInvites(driver))
	}
	return response, nil
}

func buildSchoolListInvites(drivers entity.Driver) value.SchoolListInvite {
	return value.SchoolListInvite{
		ID:           drivers.ID,
		Email:        drivers.Email,
		Name:         drivers.Name,
		Phone:        drivers.Phone,
		ProfileImage: drivers.ProfileImage,
	}
}
