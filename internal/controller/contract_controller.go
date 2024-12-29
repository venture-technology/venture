package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
)

type ContractController struct {
}

func NewContractController() *ContractController {
	return &ContractController{}
}

func (coh *ContractController) Create(c *gin.Context) {
	var input entity.Contract

	if err := c.BindJSON(&input); err != nil {
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

func (coh *ContractController) Get(c *gin.Context) {
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

func (coh *ContractController) FindAllByRg(c *gin.Context) {
	rg := c.Param("rg")

	contracts, err := coh.contractUseCase.FindAllByRg(c, &rg)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "contrato não encontrado"))
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (coh *ContractController) FindAllByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	contracts, err := coh.contractUseCase.FindAllByCnpj(c, &cnpj)

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
