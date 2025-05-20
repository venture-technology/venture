package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/controller"
)

type AdminControllers struct {
	Event      *controller.EventController
	Integrator *controller.IntegratorController
	Admin      *controller.AdminController
}

func NewAdminController() *AdminControllers {
	return &AdminControllers{
		Event:      controller.NewEventController(),
		Integrator: controller.NewIntegratorController(),
		Admin:      controller.NewAdminController(),
	}
}

func (route *AdminControllers) AdminRoutes(group *gin.RouterGroup) {

	// Event routes
	group.POST("/events", route.Event.PostV1CreateEvent)
	group.GET("/events", route.Event.GetV1ListEvents)
	group.PATCH("/events", route.Event.PatchV1UpdateEvent)
	group.DELETE("/events", route.Event.DeleteV1DeleteEvent)
	group.GET("/integrator/:id/events", route.Event.GetV1ListEventsByIntegrator)

	// Admin routes
	group.POST("/admin", route.Admin.PostV1CreateAdmin)
	group.PATCH("/admin", route.Admin.PatchV1UpdateAdmin)
	group.DELETE("/admin", route.Admin.DeleteV1DeleteAdmin)
	group.POST("/admin/login", route.Admin.PostV1LoginAdmin)

	// Integrator routes
	group.POST("/integrators", route.Integrator.PostV1CreateIntegrator)
	group.GET("/integrators", route.Integrator.GetV1ListIntegrators)
	group.PATCH("/integrators", route.Integrator.PatchV1UpdateIntegrator)
	group.DELETE("/integrators", route.Integrator.DeleteV1DeleteIntegrator)
}
