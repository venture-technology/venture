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

func (tc *TemporaryContractController) GetV1ResponsibleTempContracts(c *gin.Context) {
	cpf := c.Param("cpf")

	usecase := usecase.NewGetTempContractsResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	contracts, err := usecase.GetResponsibleTempContracts(cpf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar contratos"})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (tc *TemporaryContractController) GetV1DriverTempContracts(c *gin.Context) {
	cnh := c.Param("cnh")

	usecase := usecase.NewGetTempContractsDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	contracts, err := usecase.GetDriverTempContracts(cnh)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar contratos"})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (tc *TemporaryContractController) PostV1CancelTempContracts(httpContext *gin.Context) {
	uuid := httpContext.Param("uuid")

	usecase := usecase.NewCancelTempContractUsecase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.CancelTempContract(uuid)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar contratos"})
		return
	}

	httpContext.JSON(http.StatusOK, http.NoBody)
}
