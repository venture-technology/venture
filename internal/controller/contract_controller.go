package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/adapter"
	"github.com/venture-technology/venture/internal/domain/payments"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type ContractController struct {
}

func NewContractController() *ContractController {
	return &ContractController{}
}

func (coh *ContractController) PostV1Create(c *gin.Context) {
	var input entity.Contract
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewCreateContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		adapter.NewGoogleAdapter(),
		payments.NewStripeContract(),
	)

	err := usecase.CreateContract(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao realizar a criação do contrato"))
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (coh *ContractController) GetV1GetContract(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewGetContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		payments.NewStripeContract(),
	)

	contract, err := usecase.GetContract(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (coh *ContractController) GetV1ListContractSchool(c *gin.Context) {
	cnpj := c.Param("cnpj")

	usecase := usecase.NewListSchoolContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	contracts, err := usecase.ListSchoolContract(&cnpj)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractController) FindAllByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	contracts, err := coh.contractUseCase.FindAllByCpf(c, &cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractController) FindAllByCnh(c *gin.Context) {
	cnh := c.Param("cnh")

	contracts, err := coh.contractUseCase.FindAllByCnh(c, &cnh)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractController) Cancel(c *gin.Context) {

	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	err = coh.contractUseCase.Cancel(c, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar cancelar o contrato"))
		return
	}

	c.JSON(http.StatusOK, "contrato cancelado com sucesso")
}

func (coh *ContractController) Expired(c *gin.Context) {

	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	err = coh.contractUseCase.Expired(c, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar expirar o contrato"))
		return
	}

	c.JSON(http.StatusOK, "contrato expirado com sucesso")
}
