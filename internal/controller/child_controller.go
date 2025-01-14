package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type ChildController struct {
}

func NewChildController() *ChildController {
	return &ChildController{}
}

func (ch *ChildController) PostV1CreateChild(c *gin.Context) {
	var input entity.Child
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewCreateChildUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.CreateChild(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar novo filho"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ch *ChildController) GetV1GetChild(c *gin.Context) {
	rg := c.Param("rg")

	usecase := usecase.NewGetChildUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	child, err := usecase.GetChild(&rg)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, child)
}

func (ch *ChildController) GetV1ListChildren(c *gin.Context) {
	cpf := c.Param("cpf")

	usecase := usecase.NewListChildrenUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	children, err := usecase.ListChildren(&cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, children)
}

func (ch *ChildController) PatchV1UpdateController(c *gin.Context) {
	rg := c.Param("rg")
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewUpdateChildUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.UpdateChild(rg, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno ao atualizar informações"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (ch *ChildController) DeleteV1DeleteChild(c *gin.Context) {
	rg := c.Param("rg")

	usecase := usecase.NewDeleteChildUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.DeleteChild(&rg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar filho"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
