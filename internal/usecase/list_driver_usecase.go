package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type ListDriverFromSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListDriverFromSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListDriverFromSchoolUseCase {
	return &ListDriverFromSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (gduc *ListDriverFromSchoolUseCase) ListDriverFromSchool(cnpj string) ([]value.ListDriver, error) {
	drivers, err := gduc.repositories.PartnerRepository.FindAllByCnpj(cnpj)
	if err != nil {
		return []value.ListDriver{}, err
	}
	response := []value.ListDriver{}
	for _, driver := range drivers {
		response = append(response, buildListDriver(driver))
	}
	return response, nil
}

func buildListDriver(partner entity.Partner) value.ListDriver {
	return value.ListDriver{
		ID:            partner.Driver.ID,
		Name:          partner.Driver.Name,
		Email:         partner.Driver.Email,
		Phone:         partner.Driver.Phone,
		ProfileImage:  partner.Driver.ProfileImage,
		Accessibility: partner.Driver.Accessibility,
		States:        partner.Driver.States,
		City:          partner.Driver.City,
	}
}
