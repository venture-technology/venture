package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
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

func (lsiuc *ListSchoolInvitesUseCase) ListSchoolInvites() {

}
