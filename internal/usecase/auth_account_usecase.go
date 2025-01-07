package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type AuthAccountUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewAuthAccountUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *AuthAccountUsecase {
	return &AuthAccountUsecase{
		repositories: repositories,
		logger:       logger,
	}
}

// FindToAuth is a method from AuthAccountUsecase struct, it can be return responsible, driver or school
func (aauc *AuthAccountUsecase) FindToAuth(auth value.AuthParams) (interface{}, error) {
	var kinds = map[string]func(value.AuthParams) (interface{}, error){
		"responsible": func(authParams value.AuthParams) (interface{}, error) {
			return aauc.repositories.ResponsibleRepository.FindByEmail(authParams.Email)
		},
		"driver": func(authParams value.AuthParams) (interface{}, error) {
			return aauc.repositories.DriverRepository.FindByEmail(authParams.Email)
		},
		"school": func(authParams value.AuthParams) (interface{}, error) {
			return aauc.repositories.SchoolRepository.FindByEmail(authParams.Email)
		},
	}

	if kind, ok := kinds[auth.Kind]; ok {
		return kind(auth)
	}
	return nil, fmt.Errorf("kind not found")
}
