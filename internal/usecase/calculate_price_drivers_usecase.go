package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/domain/service/address"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type CalculatePriceDriversUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	address      address.Address
}

func NewCalculatePriceDriversUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	address address.Address,
) *CalculatePriceDriversUseCase {
	return &CalculatePriceDriversUseCase{
		repositories: repositories,
		logger:       logger,
		address:      address,
	}
}

func (cpuc *CalculatePriceDriversUseCase) CalculatePrice(
	responsibleCPF, schoolCNPJ string,
) ([]value.ListDriverToCalcPrice, error) {

	responsible, err := cpuc.repositories.ResponsibleRepository.Get(responsibleCPF)
	if err != nil {
		cpuc.logger.Infof(fmt.Sprintf("error getting responsible: %v", err))
		return nil, err
	}

	school, err := cpuc.repositories.SchoolRepository.Get(schoolCNPJ)
	if err != nil {
		cpuc.logger.Infof(fmt.Sprintf("error getting school: %v", err))
		return nil, err
	}

	distance, err := cpuc.address.Distance(
		responsible.Address.GetFullAddress(),
		school.Address.GetFullAddress(),
	)

	drivers, err := cpuc.repositories.PartnerRepository.GetBySchool(schoolCNPJ)
	if err != nil {
		cpuc.logger.Infof(fmt.Sprintf("error getting drivers: %v", err))
		return nil, err
	}

	response := []value.ListDriverToCalcPrice{}
	for _, driver := range drivers {
		if driver.Driver.Seats.Remaining > 0 {
			response = append(response, buildListDriverWithAmount(driver, *distance))
		}
	}
	return response, nil

}

func buildListDriverWithAmount(partner entity.Partner, distance float64) value.ListDriverToCalcPrice {
	return value.ListDriverToCalcPrice{
		ID:           partner.Driver.ID,
		Name:         partner.Driver.Name,
		Email:        partner.Driver.Email,
		Phone:        partner.Driver.Phone,
		ProfileImage: partner.Driver.ProfileImage,
		Amount:       partner.Driver.Amount,
		PriceTotal:   utils.CalculateContract(distance, partner.Driver.Amount),
	}
}
