package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/controller"
)

type V1Controllers struct {
	Responsible       *controller.ResponsibleController
	Kid               *controller.KidController
	School            *controller.SchoolController
	Driver            *controller.DriverController
	Invite            *controller.InviteController
	Partner           *controller.PartnerController
	Contract          *controller.ContractController
	Price             *controller.PriceController
	Webhook           *controller.WebhookController
	TemporaryContract *controller.TemporaryContractController
}

func NewV1Controller() *V1Controllers {
	return &V1Controllers{
		Invite:            controller.NewInviteController(),
		Responsible:       controller.NewResponsibleController(),
		Kid:               controller.NewKidController(),
		School:            controller.NewSchoolController(),
		Driver:            controller.NewDriverController(),
		Partner:           controller.NewPartnerController(),
		Contract:          controller.NewContractController(),
		Price:             controller.NewPriceController(),
		Webhook:           controller.NewWebhookController(),
		TemporaryContract: controller.NewTemporaryContractController(),
	}
}

func (route *V1Controllers) V1Routes(group *gin.RouterGroup) {
	group.POST("/responsible", route.Responsible.PostV1CreateResponsible)
	group.GET("/responsible/:cpf", route.Responsible.GetV1GetResponsible)
	group.PATCH("/responsible/:cpf", route.Responsible.PatchV1UpdateResponsible)
	group.DELETE("/responsible/:cpf", route.Responsible.DeleteV1DeleteResponsbile)

	group.POST("/kid", route.Kid.PostV1CreateKid)
	group.GET("/kid/:rg", route.Kid.GetV1GetKid)
	group.GET("/:cpf/kid", route.Kid.GetV1ListKids)
	group.PATCH("/kid/:rg", route.Kid.PatchV1UpdateController)
	group.DELETE("/kid/:rg", route.Kid.DeleteV1DeleteKid)

	group.POST("/school", route.School.PostV1CreateSchool)
	group.GET("/school", route.School.GetV1ListSchool)
	group.GET("/school/:cnpj", route.School.GetV1GetSchool)
	group.PATCH("/school/:cnpj", route.School.PatchV1UpdateSchool)
	group.DELETE("/school/:cnpj", route.School.DeleteV1DeleteSchool)

	group.POST("/driver", route.Driver.PostV1Create)
	group.GET("/driver/:cnh", route.Driver.GetV1GetDriver)
	group.PATCH("/driver/:cnh", route.Driver.PatchV1UpdateDriver)
	group.DELETE("/driver/:cnh", route.Driver.DeleteV1DeleteDriver)

	group.POST("/invite", route.Invite.PostV1SendInvite)
	group.GET("/driver/invite/:cnh", route.Invite.GetV1DriverListInvite)
	group.GET("/school/invite/:cnpj", route.Invite.GetV1SchoolListInvite)
	group.PATCH("/invite/:id/accept", route.Invite.PatchV1AcceptInvite)
	group.DELETE("/invite/:id/decline", route.Invite.DeleteV1DeclineInvite)

	group.GET("/driver/partner/:cnh", route.Partner.GetV1DriverListPartners)
	group.GET("/school/partner/:cnpj", route.Partner.GetV1SchoolListPartners)
	group.DELETE("/partner/:id", route.Partner.DeleteV1DeletePartner)

	group.POST("/contract", route.Contract.PostV1CreateContract)
	group.GET("/contract/:id", route.Contract.GetV1GetContract)
	group.GET("/driver/contract/:cnh", route.Contract.GetV1ListDriverContract)
	group.GET("/school/contract/:cnpj", route.Contract.GetV1ListContractSchool)
	group.GET("/responsible/contract/:cpf", route.Contract.GetV1ListResponsibleContract)
	group.PATCH("/contract/:id/cancel", route.Contract.PatchV1CancelContract)
	group.PATCH("/contract/:id/expired", route.Contract.PatchV1ExpiredContract)

	group.POST("/webhook/signature/events", route.Webhook.PostV1WebhookEvents)

	group.GET("/price/:cpf/:cnpj", route.Price.GetV1PriceDriver)

	group.GET("/temp_contracts/:cpf", route.TemporaryContract.GetV1TempContracts)
}
