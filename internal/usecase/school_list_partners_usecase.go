package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type SchoolListPartnersUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewSchoolListPartnersUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *SchoolListPartnersUseCase {
	return &SchoolListPartnersUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (slpuc *SchoolListPartnersUseCase) SchoolListPartners() {

}

func buildSchoolListPartner(partners entity.Partner) value.SchoolListPartners {
	return value.SchoolListPartners{
		ID:     partners.ID,
		Name:   partners.Driver.Name,
		Email:  partners.Driver.Email,
		Phone:  partners.Driver.Phone,
		QrCode: partners.Driver.QrCode,
		Car: fmt.Sprintf(
			"%s %s %s",
			partners.Driver.Car.Model,
			partners.Driver.Car.Year,
		),
		ProfileImage: partners.Driver.ProfileImage,
	}
}
