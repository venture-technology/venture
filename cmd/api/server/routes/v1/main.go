package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/controller"
)

type V1Controllers struct {
	Responsible controller.ResponsibleController
	Child       controller.ChildController
	School      controller.SchoolController
	Driver      controller.DriverController
	Invite      controller.InviteController
	Partner     controller.PartnerController
	Contract    controller.ContractController
	Maps        controller.MapsController
	Auth        controller.AuthController
}

func NewV1Controller() *V1Controllers {
	return &V1Controllers{
		Responsible: controller.NewResponsibleController(),
		Child:       controller.NewChildController(),
		School:      controller.NewSchoolController(),
		Driver:      controller.NewDriverController(),
		Invite:      controller.NewInviteController(),
		Partner:     controller.NewPartnerController(),
		Contract:    controller.NewContractController(),
		Maps:        controller.NewMapsController(),
		Auth:        controller.NewAuthController(),
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
	group.GET("/driver/:cnh/gallery", route.Driver.GetGallery)
	group.POST("/invite", route.Invite.Create)
	group.GET("/invite/:id", route.Invite.Get)
	group.GET("/driver/invite/:cnh", route.Invite.FindAllByCnh)
	group.GET("/school/invite/:cnpj", route.Invite.FindAllByCnpj)
	group.PATCH("/invite/:id/accept", route.Invite.Accept)
	group.DELETE("/invite/:id/decline", route.Invite.Decline)
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
	group.POST("/school/auth", route.Auth.AuthSchool)
	group.POST("/driver/auth", route.Auth.AuthDriver)
	group.POST("/responsible/auth", route.Auth.AuthResponsible)
}
