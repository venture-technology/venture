package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ah *AuthController) AuthSchool(c *gin.Context) {
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

func (ah *AuthController) AuthDriver(c *gin.Context) {
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

func (ah *AuthController) AuthResponsible(c *gin.Context) {
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
