package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type ResponsibleLoginUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewResponsibleLoginUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ResponsibleLoginUsecase {
	return &ResponsibleLoginUsecase{
		repositories: repositories,
		logger:       logger,
	}
}

func (rluc *ResponsibleLoginUsecase) LoginResponsible(email, password string) (string, error) {
	responsible, err := rluc.repositories.ResponsibleRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = utils.ValidateHash(responsible.Password, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"responsible": responsible,
		"role":        "responsible",
		"exp":         time.Now().Add(time.Hour * 240).Unix(),
	})

	return token.SignedString([]byte(value.GetJWTSecret()))
}
