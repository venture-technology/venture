package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/auth"
)

type AuthHandler struct {
	authUseCase *auth.AuthUseCase
}

func NewAuthHandler(au *auth.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: au,
	}
}

func (ah *AuthHandler) AuthSchool(c *gin.Context) {
	var input entity.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	token, err := ah.authUseCase.LoginSchool(c, &input.Email)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) AuthDriver(c *gin.Context) {
	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	token, err := ah.authUseCase.LoginDriver(c, &input.Email)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) AuthResponsible(c *gin.Context) {
	var input entity.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	token, err := ah.authUseCase.LoginResponsible(c, &input.Email)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
