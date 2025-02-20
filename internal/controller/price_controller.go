package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type PriceController struct {
}

func NewPriceController() *PriceController {
	return &PriceController{}
}

func (pc *PriceController) GetV1PriceDriver(c *gin.Context) {
	responsible := c.Param("cpf")
	school := c.Param("cnpj")

	usecase := usecase.NewCalculatePriceDriversUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	response, err := usecase.CalculatePrice(responsible, school)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "calculate pricing error"})
		return
	}

	c.JSON(http.StatusOK, response)
}
