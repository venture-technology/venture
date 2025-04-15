package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type DriverLoginUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDriverLoginUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DriverLoginUsecase {
	return &DriverLoginUsecase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dluc *DriverLoginUsecase) LoginDriver(email, password string) (string, error) {
	driver, err := dluc.repositories.DriverRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = utils.ValidateHash(driver.Password, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"driver": driver,
		"role":   "driver",
		"exp":    time.Now().Add(time.Hour * 240).Unix(),
	})

	return token.SignedString([]byte(value.GetJWTSecret()))
}
