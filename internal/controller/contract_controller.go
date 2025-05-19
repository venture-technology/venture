package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
	"github.com/venture-technology/venture/internal/value"
)

type ContractController struct {
}

func NewContractController() *ContractController {
	return &ContractController{}
}

// @Summary Cria um novo contrato
// @Description Cria um novo contrato com os dados fornecidos
// @Tags Contracts
// @Accept json
// @Produce json
// @Param contract body value.CreateContractParams true "Dados do contrato"
// @Success 201 {object} agreements.ContractRequest
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /contract [post]
func (coh *ContractController) PostV1CreateContract(httpContext *gin.Context) {
	var requestParams value.CreateContractParams
	if err := httpContext.BindJSON(&requestParams); err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	infra.App.Logger.Infof(fmt.Sprintf("requestParams: %v", requestParams))
	usecase := usecase.NewCreateContractUsecase(
		&infra.App.Repositories,
		infra.App.Workers,
		infra.App.Converters,
		infra.App.Adapters,
		infra.App.Bucket,
		infra.App.Logger,
	)

	err := usecase.CreateContract(httpContext, &requestParams)
	if err != nil {
		httpContext.JSON(
			http.StatusBadRequest,
			exceptions.InternalServerResponseError(err, "erro ao tentar criar o contrato"),
		)
		return
	}

	httpContext.JSON(http.StatusCreated, http.NoBody)
}

// @Summary Busca um contrato
// @Description Retorna um contrato pelo seu UUID
// @Tags Contracts
// @Produce json
// @Param id path string true "UUID do contrato"
// @Success 200 {object} value.GetContract
// @Failure 400 {object} map[string]string
// @Router /contract/{id} [get]
func (coh *ContractController) GetV1GetContract(httpContext *gin.Context) {
	id := httpContext.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewGetContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	contract, err := usecase.GetContract(uuid)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	httpContext.JSON(http.StatusOK, contract)
}

// @Summary Lista contratos de uma escola
// @Description Retorna todos os contratos associados ao CNPJ da escola
// @Tags Contracts
// @Produce json
// @Param cnpj path string true "CNPJ da escola"
// @Success 200 {array} []value.SchoolListContracts
// @Failure 400 {object} map[string]string
// @Router /contract/school/{cnpj} [get]
func (coh *ContractController) GetV1ListContractSchool(httpContext *gin.Context) {
	cnpj := httpContext.Param("cnpj")

	usecase := usecase.NewListSchoolContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	contracts, err := usecase.ListSchoolContract(cnpj)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	httpContext.JSON(http.StatusOK, contracts)
}

// @Summary Lista contratos de um responsável
// @Description Retorna todos os contratos associados ao CPF do responsável
// @Tags Contracts
// @Produce json
// @Param cpf path string true "CPF do responsável"
// @Success 200 {array} []value.ResponsibleListContracts
// @Failure 400 {object} map[string]string
// @Router /contract/responsible/{cpf} [get]
func (coh *ContractController) GetV1ListResponsibleContract(httpContext *gin.Context) {
	cpf := httpContext.Param("cpf")

	usecase := usecase.NewListResponsibleContractsUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	contracts, err := usecase.ListResponsibleContracts(cpf)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	httpContext.JSON(http.StatusOK, contracts)
}

// @Summary Lista contratos de um motorista
// @Description Retorna todos os contratos associados ao CNH do motorista
// @Tags Contracts
// @Produce json
// @Param cnh path string true "CNH do motorista"
// @Success 200 {array} []value.DriverListContracts
// @Failure 400 {object} map[string]string
// @Router /contract/driver/{cnh} [get]
func (coh *ContractController) GetV1ListDriverContract(httpContext *gin.Context) {
	cnh := httpContext.Param("cnh")

	usecase := usecase.NewListDriverContractsUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	contracts, err := usecase.ListDriverContracts(cnh)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	httpContext.JSON(http.StatusOK, contracts)
}

// @Summary Cancela um contrato
// @Description Cancela um contrato pelo UUID
// @Tags Contracts
// @Produce json
// @Param id path string true "UUID do contrato"
// @Success 200 {string} string "contrato cancelado com sucesso"
// @Failure 400 {object} map[string]string
// @Router /contract/{id} [post]
func (coh *ContractController) PostV1CancelContract(httpContext *gin.Context) {
	id := httpContext.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewCancelContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	err = usecase.CancelContract(uuid)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar cancelar o contrato"))
		return
	}

	httpContext.JSON(http.StatusOK, "contrato cancelado com sucesso")
}
