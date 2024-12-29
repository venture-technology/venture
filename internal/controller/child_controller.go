package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
)

type ChildController struct {
}

func NewChildController() *ChildController {
	return &ChildController{}
}

func (ch *ChildController) Create(c *gin.Context) {
	var input entity.Child

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	err := ch.childUseCase.Create(c, &input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar novo filho"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ch *ChildController) Get(c *gin.Context) {
	rg := c.Param("rg")

	child, err := ch.childUseCase.Get(c, &rg)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, child)
}

func (ch *ChildController) FindAll(c *gin.Context) {
	cpf := c.Param("cpf")

	children, err := ch.childUseCase.FindAll(c, &cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, children)
}

func (ch *ChildController) Update(c *gin.Context) {
	rg := c.Param("rg")

	var input entity.Child

	input.RG = rg

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	err := ch.childUseCase.Update(c, &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno ao atualizar informações"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (ch *ChildController) Delete(c *gin.Context) {
	rg := c.Param("rg")

	err := ch.childUseCase.Delete(c, &rg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar filho"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
