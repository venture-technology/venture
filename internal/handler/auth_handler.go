package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/auth"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authUseCase *auth.AuthUseCase
	logger      *zap.Logger
}

func NewAuthHandler(au *auth.AuthUseCase, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authUseCase: au,
		logger:      logger,
	}
}

func (ah *AuthHandler) AuthSchool(c *gin.Context) {
	var input entity.School

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	token, err := ah.authUseCase.LoginSchool(c, &input.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "escola não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) AuthDriver(c *gin.Context) {
	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	token, err := ah.authUseCase.LoginDriver(c, &input.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "motorista não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) AuthResponsible(c *gin.Context) {
	var input entity.Responsible

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	token, err := ah.authUseCase.LoginResponsible(c, &input.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "responsável não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
