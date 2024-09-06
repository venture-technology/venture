package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/usecase/contract"
)

type ContractHandler struct {
	contractUseCase *contract.ContractUseCase
}

func NewContractHandler(cu *contract.ContractUseCase) *ContractHandler {
	return &ContractHandler{
		contractUseCase: cu,
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
		log.Printf("error to create contract: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when creating contract"))
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (coh *ContractHandler) Get(c *gin.Context) {

}

func (coh *ContractHandler) FindAllByCnpj(c *gin.Context) {

}

func (coh *ContractHandler) FindAllByCpf(c *gin.Context) {

}

func (coh *ContractHandler) FindAllByCnh(c *gin.Context) {

}

func (coh *ContractHandler) Cancel(c *gin.Context) {

}

func (coh *ContractHandler) Expired(c *gin.Context) {

}
