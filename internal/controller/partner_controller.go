package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/service"
)

type PartnerController struct {
	partnerservice *service.PartnerService
}

func NewPartnerController(partnerservice *service.PartnerService) *PartnerController {
	return &PartnerController{
		partnerservice: partnerservice,
	}
}

func (ct *PartnerController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("api/v1/partner")

	api.GET("/driver", ct.GetPartnersByDriver) // para que o motorista busque todos seus parceiros
	api.GET("/school", ct.GetPartnersBySchool) // para que a escola busque todos seus parceiros
	api.DELETE("/:cnh", ct.DeletePartner)      // para que a escola exclua um parceiro da sua rede
}

func (ct *PartnerController) GetPartnersByDriver(c *gin.Context) {

}

func (ct *PartnerController) GetPartnersBySchool(c *gin.Context) {

}

func (ct *PartnerController) DeletePartner(c *gin.Context) {

}
