package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/usecase/contract"
	"go.uber.org/zap"
)

type ContractHandler struct {
	contractUseCase *contract.ContractUseCase
	logger          *zap.Logger
}

func NewContractHandler(
	cu *contract.ContractUseCase,
	logger *zap.Logger,
) *ContractHandler {
	return &ContractHandler{
		contractUseCase: cu,
		logger:          logger,
	}
}

func (coh *ContractHandler) Create(c *gin.Context) {
	var input entity.Contract

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	err := coh.contractUseCase.Create(c, &input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao realizar a criação do contrato"))
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (coh *ContractHandler) Get(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	contract, err := coh.contractUseCase.Get(c, uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (coh *ContractHandler) FindAllByRg(c *gin.Context) {
	rg := c.Param("rg")

	contracts, err := coh.contractUseCase.FindAllByRg(c, &rg)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractHandler) FindAllByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	contracts, err := coh.contractUseCase.FindAllByCnpj(c, &cnpj)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractHandler) FindAllByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	contracts, err := coh.contractUseCase.FindAllByCpf(c, &cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractHandler) FindAllByCnh(c *gin.Context) {
	cnh := c.Param("cnh")

	contracts, err := coh.contractUseCase.FindAllByCnh(c, &cnh)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractHandler) Cancel(c *gin.Context) {

	id := c.Param("id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		log.Printf("error to parse id: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	err = coh.contractUseCase.Cancel(c, uuid)

	if err != nil {
		log.Printf("error while cancel contract: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, "contrato cancelado com sucesso")

}

func (coh *ContractHandler) Expired(c *gin.Context) {

}
