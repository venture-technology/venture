package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type TemporaryContractController struct {
}

func NewTemporaryContractController() *TemporaryContractController {
	return &TemporaryContractController{}
}

// @Summary Busca contratos temporários de um responsável
// @Description Retorna todos os contratos temporários associados ao CPF do responsável
// @Tags Temporary Contracts
// @Produce json
// @Param cpf path string true "CPF do responsável"
// @Success 200 {array} []value.GetTempContracts
// @Failure 500 {object} map[string]string
// @Router /temporary-contract/responsible/{cpf} [get]
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

// @Summary Busca contratos temporários de um motorista
// @Description Retorna todos os contratos temporários associados ao CNH do motorista
// @Tags Temporary Contracts
// @Produce json
// @Param cnh path string true "CNH do motorista"
// @Success 200 {array} []value.GetTempContracts
// @Failure 500 {object} map[string]string
// @Router /temporary-contract/driver/{cnh} [get]
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

// @Summary Cancela um contrato temporário
// @Description Cancela um contrato temporário pelo UUID
// @Tags Temporary Contracts
// @Produce json
// @Param uuid path string true "UUID do contrato temporário"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /temporary-contract/{uuid} [post]
func (tc *TemporaryContractController) PostV1CancelTempContracts(httpContext *gin.Context) {
	uuid, err := uuid.Parse(httpContext.Param("uuid"))
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "uuid inválido"})
		return
	}

	usecase := usecase.NewCancelTempContractUsecase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	if err := usecase.CancelTempContract(uuid); err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao cancelar contrato"})
		return
	}

	httpContext.JSON(http.StatusOK, http.NoBody)
}
