package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/controller"
	"github.com/venture-technology/venture/internal/domain/service/middleware"
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
	// middlewares
	// responsible, driver, school
	rm, dm, sm := middlewares()

	// responsible
	group.POST("/responsible", route.Responsible.PostV1CreateResponsible)
	group.GET("/responsible/:cpf", route.Responsible.GetV1GetResponsible)
	group.PATCH("/responsible/:cpf", rm.Middleware(), route.Responsible.PatchV1UpdateResponsible)
	group.DELETE("/responsible/:cpf", rm.Middleware(), route.Responsible.DeleteV1DeleteResponsbile)
	group.POST("/responsible/login", route.Responsible.PostV1LoginResponsible)

	// kid
	group.POST("/kid/:cpf", rm.Middleware(), route.Kid.PostV1CreateKid)
	group.GET("/kid/:rg", route.Kid.GetV1GetKid)
	group.GET("/kids/:cpf", rm.Middleware(), route.Kid.GetV1ListKids)
	group.PATCH("/kid/:cpf/:rg", rm.Middleware(), route.Kid.PatchV1UpdateController)
	group.DELETE("/kid/:cpf/:rg", rm.Middleware(), route.Kid.DeleteV1DeleteKid)

	// school
	group.POST("/school", route.School.PostV1CreateSchool)
	group.GET("/schools", route.School.GetV1ListSchool)
	group.GET("/school/:cnpj", route.School.GetV1GetSchool)
	group.PATCH("/school/:cnpj", sm.Middleware(), route.School.PatchV1UpdateSchool)
	group.DELETE("/school/:cnpj", sm.Middleware(), route.School.DeleteV1DeleteSchool)
	group.POST("/school/login", route.School.PostV1LoginSchool)

	// driver
	group.POST("/driver", route.Driver.PostV1Create)
	group.GET("/driver/:cnh", route.Driver.GetV1GetDriver)
	group.PATCH("/driver/:cnh", dm.Middleware(), route.Driver.PatchV1UpdateDriver)
	group.DELETE("/driver/:cnh", dm.Middleware(), route.Driver.DeleteV1DeleteDriver)
	group.POST("/driver/login", route.Driver.PostV1LoginDriver)

	// invite
	group.POST("/invite", sm.Middleware(), route.Invite.PostV1SendInvite)
	group.GET("/driver/invites/:cnh", dm.Middleware(), route.Invite.GetV1DriverListInvite)
	group.GET("/school/invites/:cnpj", sm.Middleware(), route.Invite.GetV1SchoolListInvite)
	group.PATCH("/invite/:id/accept", dm.Middleware(), route.Invite.PatchV1AcceptInvite)
	group.DELETE("/invite/:id/decline", dm.Middleware(), route.Invite.DeleteV1DeclineInvite)

	// partner
	group.GET("/driver/partners/:cnh", dm.Middleware(), route.Partner.GetV1DriverListPartners)
	group.GET("/school/partners/:cnpj", sm.Middleware(), route.Partner.GetV1SchoolListPartners)
	group.DELETE("/partner/:id", sm.Middleware(), route.Partner.DeleteV1DeletePartner)

	// contract
	group.POST("/contract", rm.Middleware(), route.Contract.PostV1CreateContract)
	group.GET("/contract/:id", route.Contract.GetV1GetContract)
	group.GET("/driver/contracts/:cnh", dm.Middleware(), route.Contract.GetV1ListDriverContract)
	group.GET("/school/contracts/:cnpj", sm.Middleware(), route.Contract.GetV1ListContractSchool)
	group.GET("/responsible/contracts/:cpf", rm.Middleware(), route.Contract.GetV1ListResponsibleContract)
	group.POST("/contract/:id/cancel", rm.Middleware(), route.Contract.PostV1CancelContract)

	// webhook
	group.POST("/webhook/signature/events", route.Webhook.PostV1WebhookSignatureEvents)
	group.POST("/webhook/payments/events", route.Webhook.PostV1WebhookPaymentsEvents)

	// price
	group.GET("/price/:cpf/:cnpj", rm.Middleware(), route.Price.GetV1PriceDriver)

	// temporary contract
	group.GET("/temporary-contract/responsible/:cpf", rm.Middleware(), route.TemporaryContract.GetV1ResponsibleTempContracts)
	group.GET("/temporary-contract/driver/:cnh", dm.Middleware(), route.TemporaryContract.GetV1DriverTempContracts)
	group.POST("/temporary-contract/cancel/:uuid", rm.Middleware(), route.TemporaryContract.PostV1CancelTempContracts)
}

func middlewares() (
	*middleware.ResponsibleMiddleware,
	*middleware.DriverMiddleware,
	*middleware.SchoolMiddleware,
) {
	return middleware.NewResponsibleMiddleware(),
		middleware.NewDriverMiddleware(),
		middleware.NewSchoolMiddleware()
}
