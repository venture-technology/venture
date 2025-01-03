package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/controller"
)

type V1Controllers struct {
	Responsible *controller.ResponsibleController
	Child       *controller.ChildController
	School      *controller.SchoolController
	Driver      *controller.DriverController
	Invite      *controller.InviteController
	Partner     *controller.PartnerController
	Contract    *controller.ContractController
	Maps        *controller.MapsController
}

func NewV1Controller() *V1Controllers {
	return &V1Controllers{
		Invite:      controller.NewInviteController(),
		Responsible: controller.NewResponsibleController(),
		Child:       controller.NewChildController(),
		School:      controller.NewSchoolController(),
		Driver:      controller.NewDriverController(),
		Partner:     controller.NewPartnerController(),
		Contract:    controller.NewContractController(),
		Maps:        controller.NewMapsController(),
	}
}

func (route *V1Controllers) V1Routes(group *gin.RouterGroup) {
	group.POST("/responsible", route.Responsible.Create)
	group.GET("/responsible/:cpf", route.Responsible.Get)
	group.PATCH("/responsible/:cpf", route.Responsible.Update)
	group.DELETE("/responsible/:cpf", route.Responsible.Delete)

	group.POST("/child", route.Child.Create)
	group.GET("/child/:rg", route.Child.Get)
	group.GET("/:cpf/child", route.Child.FindAll)
	group.PATCH("/child/:rg", route.Child.Update)
	group.DELETE("/child/:rg", route.Child.Delete)

	group.POST("/school", route.School.Create)
	group.GET("/school", route.School.FindAll)
	group.GET("/school/:cnpj", route.School.Get)
	group.PATCH("/school/:cnpj", route.School.Update)
	group.DELETE("/school/:cnpj", route.School.Delete)

	group.POST("/driver", route.Driver.Create)
	group.GET("/driver/:cnh", route.Driver.Get)
	group.PATCH("/driver/:cnh", route.Driver.Update)
	group.DELETE("/driver/:cnh", route.Driver.Delete)

	group.POST("/invite", route.Invite.PostV1SendInvite)
	group.GET("/driver/invite/:cnh", route.Invite.GetV1DriverListInvite)
	group.GET("/school/invite/:cnpj", route.Invite.GetV1DriverListInvite)
	group.PATCH("/invite/:id/accept", route.Invite.PatchV1AcceptInvite)
	group.DELETE("/invite/:id/decline", route.Invite.DeleteV1DeclineInvite)

	group.GET("/partner/:id", route.Partner.Get)
	group.GET("/driver/partner/:cnh", route.Partner.FindAllByCnh)
	group.GET("/school/partner/:cnpj", route.Partner.FindAllByCnpj)
	group.DELETE("/partner/:id", route.Partner.Delete)

	group.POST("/contract", route.Contract.Create)
	group.GET("/contract/:id", route.Contract.Get)
	group.GET("/driver/contract/:cnh", route.Contract.FindAllByCnh)
	group.GET("/school/contract/:cnpj", route.Contract.FindAllByCnpj)
	group.GET("/responsible/contract/:cpf", route.Contract.FindAllByCpf)
	group.GET("/child/contract/:rg", route.Contract.FindAllByRg)
	group.PATCH("/contract/:id/cancel", route.Contract.Cancel)
	group.PATCH("/contract/:id/expired", route.Contract.Expired)

	group.POST("/maps/price", route.Maps.CalculatePrice)
}
