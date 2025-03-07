package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type TemporaryContractController struct {
}

func NewTemporaryContractController() *TemporaryContractController {
	return &TemporaryContractController{}
}

func (tc *TemporaryContractController) GetV1TempContracts(c *gin.Context) {
	cpf := c.Param("cpf")

	usecase := usecase.NewGetTempContractsUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	contracts, err := usecase.GetTempContracts(cpf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar contratos"})
		return
	}

	c.JSON(http.StatusOK, contracts)
}
