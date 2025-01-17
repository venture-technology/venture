package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type GetResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewGetResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *GetResponsibleUseCase {
	return &GetResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (gruc *GetResponsibleUseCase) GetResponsible(cpf string) (value.GetResponsible, error) {
	responsible, err := gruc.repositories.ResponsibleRepository.Get(cpf)
	if err != nil {
		return value.GetResponsible{}, err
	}
	return value.GetResponsible{
		ID:    responsible.ID,
		Name:  responsible.Name,
		Email: responsible.Email,
		Phone: responsible.Phone,
		Address: utils.BuildAddress(
			responsible.Address.Street,
			responsible.Address.Number,
			responsible.Address.Complement,
			responsible.Address.Zip,
		),
		CustomerId:      responsible.CustomerId,
		ProfileImage:    responsible.ProfileImage,
		PaymentMethodId: responsible.PaymentMethodId,
		CreatedAt:       responsible.CreatedAt,
	}, nil
}
