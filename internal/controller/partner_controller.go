package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PartnerController struct {
}

func NewPartnerController() *PartnerController {
	return &PartnerController{}
}

func (ph *PartnerController) Get(c *gin.Context) {
	id := c.Param("id")

	partner, err := ph.PartnerUseCase.Get(c, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, partner)
}

func (ph *PartnerController) FindAllByCnh(c *gin.Context) {
	cnh := c.Param("cnh")

	partners, err := ph.PartnerUseCase.FindAllByCnh(c, &cnh)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, partners)
}

func (ph *PartnerController) FindAllByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	partners, err := ph.PartnerUseCase.FindAllByCnpj(c, &cnpj)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, partners)
}

func (ph *PartnerController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := ph.PartnerUseCase.Delete(c, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
