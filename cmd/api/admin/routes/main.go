package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/controller"
)

type AdminControllers struct {
	Admin *controller.AdminController
}

func NewAdminController() *AdminControllers {
	return &AdminControllers{
		Admin: controller.NewAdminController(),
	}
}

func (route *AdminControllers) AdminRoutes(group *gin.RouterGroup) {
	group.POST("/api-key/:name", route.Admin.NewApiKey)
}
