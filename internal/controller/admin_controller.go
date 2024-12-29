package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/application"
	"github.com/venture-technology/venture/internal/infra"
)

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (ah *AdminController) NewApiKey(c *gin.Context) {
	name := c.Param("name")

	usecase := application.NewGenerateApiKeyAdminService(
		&infra.App.RedisRepositories,
		infra.App.Logger,
	)

	err := usecase.NewApiKey(c, name)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao criar nova chave"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
