package controller

import (
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (ah *AdminController) PostV1CreateAdmin(httpContext *gin.Context) {
}

func (ah *AdminController) PatchV1UpdateAdmin(httpContext *gin.Context) {
}

func (ah *AdminController) DeleteV1DeleteAdmin(httpContext *gin.Context) {
}

func (ah *AdminController) PostV1LoginAdmin(httpContext *gin.Context) {
}
