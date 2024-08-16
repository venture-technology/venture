package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/usecase/responsible"
)

type ResponsibleHandler struct {
	responsibleUseCase *responsible.ResponsibleUseCase
}

func NewResponsibleHandler(responsibleUseCase *responsible.ResponsibleUseCase) *ResponsibleHandler {
	return &ResponsibleHandler{
		responsibleUseCase: responsibleUseCase,
	}
}

func (rh *ResponsibleHandler) Create(c *gin.Context) {

}

func (rh *ResponsibleHandler) Get(c *gin.Context) {

}

func (rh *ResponsibleHandler) Update(c *gin.Context) {

}

func (rh *ResponsibleHandler) Delete(c *gin.Context) {

}

func (rh *ResponsibleHandler) SaveCard(c *gin.Context) {

}
