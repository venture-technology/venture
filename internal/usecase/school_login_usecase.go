package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type SchoolLoginUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewSchoolLoginUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *SchoolLoginUsecase {
	return &SchoolLoginUsecase{
		repositories: repositories,
		logger:       logger,
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
		"role":   "school",
		"exp":    time.Now().Add(time.Hour * 240).Unix(),
	})

	return token.SignedString([]byte(value.GetJWTSecret()))
}
