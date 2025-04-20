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

// @Summary Calcula o preço dos motoristas
// @Description Retorna os preços estimados com base no responsável e na escola
// @Tags Prices
// @Produce json
// @Param cpf path string true "CPF do responsável"
// @Param cnpj path string true "CNPJ da escola"
// @Success 200 {object} value.ListDriverToCalcPrice
// @Failure 400 {object} map[string]string
// @Router /price/{cpf}/{cnpj} [get]
func (pc *PriceController) GetV1PriceDriver(httpContext *gin.Context) {
	responsible := httpContext.Param("cpf")
	school := httpContext.Param("cnpj")

	usecase := usecase.NewCalculatePriceDriversUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	response, err := usecase.CalculatePrice(responsible, school)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "calculate pricing error"})
		return
	}

	httpContext.JSON(http.StatusOK, response)
}
