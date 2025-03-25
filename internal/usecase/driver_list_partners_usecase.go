package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
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
	partners, err := dlpuc.repositories.PartnerRepository.GetByDriver(cnh)
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
		ID:    partners.Record,
		Name:  partners.School.Name,
		Email: partners.School.Email,
		Phone: partners.School.Phone,
		Address: utils.BuildAddress(
			partners.School.Address.Street,
			partners.School.Address.Number,
			partners.School.Address.Complement,
			partners.School.Address.Zip,
		),
		ProfileImage: partners.School.ProfileImage,
		CreatedAt:    partners.CreatedAt,
	}
}
