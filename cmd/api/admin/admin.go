package admin

import (
	"github.com/venture-technology/venture/internal/handler"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase/admin"
	"github.com/venture-technology/venture/internal/usecase/auth/middleware"
)

type Admin struct {
	app   *infra.Application
	admin *handler.AdminHandler
}

func NewAdmin(app *infra.Application) *Admin {
	return &Admin{
		app: app,
	}
}

func (adm *Admin) Setup() {

	adm.app.Adm.Use(middleware.AdminMiddleware())

	adm.admin = handler.NewAdminHandler(admin.NewAdminUseCase(repository.NewAdminRepository(adm.app.Cache)))

	adm.NewRoutes()

}

func (adm *Admin) NewRoutes() {
	adm.app.Adm.POST("/api-key/:name", adm.admin.NewApiKey)
}
