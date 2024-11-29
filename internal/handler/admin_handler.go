package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/usecase/admin"
)

type AdminHandler struct {
	adminUseCase *admin.AdminUseCase
}

func NewAdminHandler(adminUseCase *admin.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminUseCase,
	}
}

func (ah *AdminHandler) NewApiKey(c *gin.Context) {
	name := c.Param("name")

	err := ah.adminUseCase.NewApiKey(c, name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao criar nova chave"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
