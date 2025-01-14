package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type DriverListPartnersUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDriverListPartnersUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DriverListPartnersUseCase {
	return &DriverListPartnersUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dlpuc *DriverListPartnersUseCase) DriverListPartners(cnh string) ([]value.DriverListPartners, error) {
	partners, err := dlpuc.repositories.PartnerRepository.FindAllByCnh(cnh)
	if err != nil {
		return nil, err
	}
	response := []value.DriverListPartners{}
	for _, partner := range partners {
		response = append(response, buildDriverListPartner(partner))
	}
	return response, nil
}

func buildDriverListPartner(partners entity.Partner) value.DriverListPartners {
	return value.DriverListPartners{
		ID:    partners.ID,
		Name:  partners.School.Name,
		Email: partners.School.Email,
		Phone: partners.School.Phone,
		Address: fmt.Sprintf(
			"%s, %s, %s",
			partners.School.Address.Street,
			partners.School.Address.Number,
			partners.School.Address.ZIP,
		),
		ProfileImage: partners.School.ProfileImage,
	}
}
