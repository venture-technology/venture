package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/realtime"
)

var kinds = map[string]interface{}{
	"responsible": entity.Responsible{},
	"driver":      entity.Driver{},
	"school":      entity.School{},
}

func NewToken(
	conf *config.Config,
	authParams value.AuthParams,
) (string, error) {
	usecase := usecase.NewAuthAccountUsecase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	user, err := usecase.FindToAuth(authParams)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": user,
			"exp":  realtime.Now().Add(time.Hour * 24).Unix(),
		},
	)

	return token.SignedString([]byte(conf.Server.Secret))
}

func GetKind(kind string) (interface{}, error) {
	if kind, exists := kinds[kind]; exists {
		return kind, nil
	}
	return nil, fmt.Errorf("kind not found")
}
