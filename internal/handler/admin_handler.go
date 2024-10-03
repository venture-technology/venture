package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/infra/ventracer"
	"github.com/venture-technology/venture/internal/usecase/admin"
	"go.uber.org/zap"
)

type AdminHandler struct {
	logger       *zap.Logger
	tracer       *ventracer.Ventracer
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
		log.Printf("error to create a key: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao criar nova chave"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
