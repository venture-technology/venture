package usecase

import (
	"fmt"

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
	schools, err := lsiuc.repositories.InviteRepository.GetByDriver(cnh)
	if err != nil {
		return nil, err
	}
	lsiuc.logger.Infof(fmt.Sprintf("ListDriverInvitesUseCase.ListDriverInvites: %v", schools))
	response := []value.DriverListInvite{}
	for _, school := range schools {
		response = append(response, buildDriverListInvites(school))
	}
	return response, nil
}

func buildDriverListInvites(schools entity.School) value.DriverListInvite {
	return value.DriverListInvite{
		Email:        schools.Email,
		Name:         schools.Name,
		Phone:        schools.Phone,
		ProfileImage: schools.ProfileImage,
		Address: utils.BuildAddress(
			schools.Address.Street,
			schools.Address.Number,
			schools.Address.Complement,
			schools.Address.Zip,
		),
	}
}
