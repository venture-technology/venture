package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/repository"
	"go.uber.org/zap"
)

type AuthUseCase struct {
	schoolRepository      repository.ISchoolRepository
	driverRepository      repository.IDriverRepository
	responsibleRepository repository.IResponsibleRepository
	logger                *zap.Logger
}

func NewAuthUseCase(
	schoolRepository repository.ISchoolRepository,
	driverRepository repository.IDriverRepository,
	responsibleRepository repository.IResponsibleRepository,
	logger *zap.Logger,
) *AuthUseCase {
	return &AuthUseCase{
		schoolRepository:      schoolRepository,
		driverRepository:      driverRepository,
		responsibleRepository: responsibleRepository,
		logger:                logger,
	}
}

func (au *AuthUseCase) LoginSchool(ctx context.Context, email *string) (string, error) {
	conf := config.Get()

	school, err := au.schoolRepository.FindByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"school": school,
		"exp":    time.Now().Add(time.Hour * 240).Unix(),
	})

	return token.SignedString([]byte(conf.Server.Secret))
}

func (au *AuthUseCase) LoginDriver(ctx context.Context, email *string) (string, error) {
	conf := config.Get()

	driver, err := au.driverRepository.FindByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"driver": driver,
		"exp":    time.Now().Add(time.Hour * 240).Unix(),
	})

	return token.SignedString([]byte(conf.Server.Secret))
}

func (au *AuthUseCase) LoginResponsible(ctx context.Context, email *string) (string, error) {
	conf := config.Get()

	responsible, err := au.responsibleRepository.FindByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"responsible": responsible,
		"exp":         time.Now().Add(time.Hour * 240).Unix(),
	})

	return token.SignedString([]byte(conf.Server.Secret))
}
