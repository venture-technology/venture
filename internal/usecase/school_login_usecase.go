package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/pkg/utils"
)

type SchoolLoginUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	config       config.Config
}

func NewSchoolLoginUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	config config.Config,
) *SchoolLoginUsecase {
	return &SchoolLoginUsecase{
		repositories: repositories,
		logger:       logger,
		config:       config,
	}
}

func (sluc *SchoolLoginUsecase) LoginSchool(email, password string) (string, error) {
	school, err := sluc.repositories.SchoolRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = utils.ValidateHash(school.Password, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"school": school,
		"exp":    time.Now().Add(time.Hour * 240).Unix(),
	})

	return token.SignedString([]byte(sluc.config.Server.Secret))
}
