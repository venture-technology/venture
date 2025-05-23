package controller

import "github.com/gin-gonic/gin"

type IntegratorController struct {
}

func NewIntegratorController() *IntegratorController {
	return &IntegratorController{}
}

func (ih *IntegratorController) PostV1CreateIntegrator(httpContext *gin.Context) {
}

func (ih *IntegratorController) GetV1ListIntegrators(httpContext *gin.Context) {
}

func (ih *IntegratorController) PatchV1UpdateIntegrator(httpContext *gin.Context) {
}

func (ih *IntegratorController) DeleteV1DeleteIntegrator(httpContext *gin.Context) {
}

func (ih *IntegratorController) GetV1ListEventsByIntegrator(httpContext *gin.Context) {
}
