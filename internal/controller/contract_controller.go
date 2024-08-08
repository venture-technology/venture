package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/service"
)

type ContractController struct {
	contractservice *service.ContractService
}

func NewContractController(contractservice *service.ContractService) *ContractController {
	return &ContractController{
		contractservice: contractservice,
	}
}

func (ct *ContractController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("api/v1/contract")

	api.POST("/", ct.CreateContract)
	// adicionar query param de cpf
	api.GET("/:record", ct.GetContract)                    // para verificar um contrato em especifico
	api.GET("/invoice/:record", ct.GetInvoiceFromContract) // para verificar todas as faturas daquele contrato
	api.GET("/invoice/:record/:id", ct.GetInvoice)         // para verificar uma fatura de um contrato especifico
	api.PATCH("/status", ct.UpdateStatusContract)          // para atualizar e setar contrato como cancelado (vistado apenas por webhooks da stripe)
}

func (ct *ContractController) CreateContract(c *gin.Context) {

}

func (ct *ContractController) GetContractByCpf(c *gin.Context) {

}

func (ct *ContractController) GetContract(c *gin.Context) {

}

func (ct *ContractController) GetInvoiceFromContract(c *gin.Context) {

}

func (ct *ContractController) GetInvoice(c *gin.Context) {

}

func (ct *ContractController) UpdateStatusContract(c *gin.Context) {

}
