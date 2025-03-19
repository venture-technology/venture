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

func (coh *ContractController) PostV1CreateContract(c *gin.Context) {
	var requestParams value.CreateContractRequestParams
	if err := c.BindJSON(&requestParams); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	infra.App.Logger.Infof(fmt.Sprintf("requestParams: %v", requestParams))

	usecase := usecase.NewCreateContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
		infra.App.Bucket,
	)

	response, err := usecase.CreateContract(&requestParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao realizar a criação do contrato"))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// func (coh *ContractController) GetV1GetContract(c *gin.Context) {
// 	id := c.Param("id")

// 	uuid, err := uuid.Parse(id)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
// 		return
// 	}

// 	usecase := usecase.NewGetContractUseCase(
// 		&infra.App.Repositories,
// 		infra.App.Logger,
// 		infra.App.Adapters,
// 	)

// 	contract, err := usecase.GetContract(uuid)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
// 		return
// 	}

// 	c.JSON(http.StatusOK, contract)
// }

// func (coh *ContractController) GetV1ListContractSchool(c *gin.Context) {
// 	// cnpj := c.Param("cnpj")

// 	// usecase := usecase.NewListSchoolContractUseCase(
// 	// 	&infra.App.Repositories,
// 	// 	infra.App.Logger,
// 	// )

// 	// contracts, err := usecase.ListSchoolContract(&cnpj)

// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, contracts)
// }

// func (coh *ContractController) GetV1ListResponsibleContract(c *gin.Context) {
// 	cpf := c.Param("cpf")

// 	usecase := usecase.NewListResponsibleContractsUseCase(
// 		&infra.App.Repositories,
// 		infra.App.Logger,
// 		infra.App.Adapters,
// 	)

// 	contracts, err := usecase.ListResponsibleContracts(&cpf)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
// 		return
// 	}

// 	c.JSON(http.StatusOK, contracts)
// }

// func (coh *ContractController) GetV1ListDriverContract(c *gin.Context) {
// 	cnh := c.Param("cnh")

// 	usecase := usecase.NewListDriverContractsUseCase(
// 		&infra.App.Repositories,
// 		infra.App.Logger,
// 	)

// 	contracts, err := usecase.ListDriverContracts(&cnh)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
// 		return
// 	}

// 	c.JSON(http.StatusOK, contracts)
// }

func (coh *ContractController) PostV1CancelContract(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewCancelContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	err = usecase.CancelContract(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar cancelar o contrato"))
		return
	}

	c.JSON(http.StatusOK, "contrato cancelado com sucesso")
}

// func (coh *ContractController) PatchV1ExpiredContract(c *gin.Context) {
// 	id := c.Param("id")

// 	uuid, err := uuid.Parse(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
// 		return
// 	}

// 	usecase := usecase.NewExpireContractUseCase(
// 		&infra.App.Repositories,
// 		infra.App.Logger,
// 	)

// 	err = usecase.ExpireContract(uuid)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar expirar o contrato"))
// 		return
// 	}

// 	c.JSON(http.StatusOK, "contrato expirado com sucesso")
// }
