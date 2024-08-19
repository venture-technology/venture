package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/usecase/child"
)

type ChildHandler struct {
	childUseCase *child.ChildUseCase
}

func NewChildHandler(childUseCase *child.ChildUseCase) *ChildHandler {
	return &ChildHandler{
		childUseCase: childUseCase,
	}
}

func (ch *ChildHandler) Create(c *gin.Context) {

}

func (ch *ChildHandler) Get(c *gin.Context) {

}

func (ch *ChildHandler) Update(c *gin.Context) {

}

func (ch *ChildHandler) Delete(c *gin.Context) {

}
