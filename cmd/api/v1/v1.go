package v1

import (
	"github.com/venture-technology/venture/internal/domain/payments"
	"github.com/venture-technology/venture/internal/handler"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase/auth"
	"github.com/venture-technology/venture/internal/usecase/auth/middleware"
	"github.com/venture-technology/venture/internal/usecase/child"
	"github.com/venture-technology/venture/internal/usecase/contract"
	"github.com/venture-technology/venture/internal/usecase/driver"
	"github.com/venture-technology/venture/internal/usecase/invite"
	"github.com/venture-technology/venture/internal/usecase/maps"
	"github.com/venture-technology/venture/internal/usecase/partner"
	"github.com/venture-technology/venture/internal/usecase/responsible"
	"github.com/venture-technology/venture/internal/usecase/school"
)

type V1 struct {
	app         *infra.Application
	responsible *handler.ResponsibleHandler
	child       *handler.ChildHandler
	school      *handler.SchoolHandler
	driver      *handler.DriverHandler
	invite      *handler.InviteHandler
	partner     *handler.PartnerHandler
	contract    *handler.ContractHandler
	maps        *handler.MapsHandler
	auth        *handler.AuthHandler
}

func NewV1(app *infra.Application) *V1 {
	return &V1{
		app: app,
	}
}

func (v1 *V1) Setup() {

	v1.app.V1.Use(middleware.RequestMiddleware(v1.app.Cache))

	v1.responsible = handler.NewResponsibleHandler(responsible.NewResponsibleUseCase(repository.NewResponsibleRepository(v1.app.Database)))
	v1.child = handler.NewChildHandler(child.NewChildUseCase(repository.NewChildRepository(v1.app.Database)))
	v1.school = handler.NewSchoolHandler(school.NewSchoolUseCase(repository.NewSchoolRepository(v1.app.Database)))
	v1.driver = handler.NewDriverHandler(driver.NewDriverUseCase(repository.NewDriverRepository(v1.app.Database), repository.NewAwsRepository(v1.app.Cloud)))
	v1.invite = handler.NewInviteHandler(invite.NewInviteUseCase(repository.NewInviteRepository(v1.app.Database), repository.NewPartnerRepository(v1.app.Database)))
	v1.partner = handler.NewPartnerHandler(partner.NewPartnerUseCase(repository.NewPartnerRepository(v1.app.Database)))
	v1.contract = handler.NewContractHandler(contract.NewContractUseCase(repository.NewContractRepository(v1.app.Database), repository.NewChildRepository(v1.app.Database), repository.NewDriverRepository(v1.app.Database), repository.NewSchoolRepository(v1.app.Database), payments.NewStripeUseCase()))
	v1.maps = handler.NewMapsHandler(maps.NeWMapsUseCase(repository.NewMapsRepository(v1.app.Database)))
	v1.auth = handler.NewAuthHandler(auth.NewAuthUseCase(repository.NewSchoolRepository(v1.app.Database), repository.NewDriverRepository(v1.app.Database), repository.NewResponsibleRepository(v1.app.Database)))

	v1.NewRoutes()

}

func (v1 *V1) NewRoutes() {
	v1.app.V1.POST("/responsible", v1.responsible.Create)
	v1.app.V1.POST("/responsible/card", v1.responsible.SaveCard)
	v1.app.V1.GET("/responsible/:cpf", v1.responsible.Get)
	v1.app.V1.PATCH("/responsible/:cpf", v1.responsible.Update)
	v1.app.V1.DELETE("/responsible/:cpf", v1.responsible.Delete)
	v1.app.V1.POST("/child", v1.child.Create)
	v1.app.V1.GET("/child/:rg", v1.child.Get)
	v1.app.V1.GET("/:cpf/child", v1.child.FindAll)
	v1.app.V1.PATCH("/child/:rg", v1.child.Update)
	v1.app.V1.DELETE("/child/:rg", v1.child.Delete)
	v1.app.V1.POST("/school", v1.school.Create)
	v1.app.V1.GET("/school", v1.school.FindAll)
	v1.app.V1.GET("/school/:cnpj", v1.school.Get)
	v1.app.V1.PATCH("/school/:cnpj", v1.school.Update)
	v1.app.V1.DELETE("/school/:cnpj", v1.school.Delete)
	v1.app.V1.POST("/driver", v1.driver.Create)
	v1.app.V1.GET("/driver/:cnh", v1.driver.Get)
	v1.app.V1.PATCH("/driver/:cnh", v1.driver.Update)
	v1.app.V1.POST("/driver/:cnh/pix", v1.driver.SavePix)
	v1.app.V1.POST("/driver/:cnh/bank", v1.driver.SaveBank)
	v1.app.V1.DELETE("/driver/:cnh", v1.driver.Delete)
	v1.app.V1.GET("/driver/:cnh/gallery", v1.driver.GetGallery)
	v1.app.V1.POST("/invite", v1.invite.Create)
	v1.app.V1.GET("/invite/:id", v1.invite.Get)
	v1.app.V1.GET("/driver/invite/:cnh", v1.invite.FindAllByCnh)
	v1.app.V1.GET("/school/invite/:cnpj", v1.invite.FindAllByCnpj)
	v1.app.V1.PATCH("/invite/:id/accept", v1.invite.Accept)
	v1.app.V1.DELETE("/invite/:id/decline", v1.invite.Decline)
	v1.app.V1.GET("/partner/:id", v1.partner.Get)
	v1.app.V1.GET("/driver/partner/:cnh", v1.partner.FindAllByCnh)
	v1.app.V1.GET("/school/partner/:cnpj", v1.partner.FindAllByCnpj)
	v1.app.V1.DELETE("/partner/:id", v1.partner.Delete)
	v1.app.V1.POST("/contract", v1.contract.Create)
	v1.app.V1.GET("/contract/:id", v1.contract.Get)
	v1.app.V1.GET("/driver/contract", v1.contract.FindAllByCnh)
	v1.app.V1.GET("/school/contract", v1.contract.FindAllByCnpj)
	v1.app.V1.GET("/responsible/contract", v1.contract.FindAllByCpf)
	v1.app.V1.PATCH("/contract/:id/cancel", v1.contract.Cancel)
	v1.app.V1.PATCH("/webhook/contract/:id/expired", v1.contract.Expired)
	v1.app.V1.POST("/maps/price", v1.maps.CalculatePrice)
	v1.app.V1.POST("/school/auth", v1.auth.AuthSchool)
	v1.app.V1.POST("/driver/auth", v1.auth.AuthDriver)
	v1.app.V1.POST("/responsible/auth", v1.auth.AuthResponsible)
}
