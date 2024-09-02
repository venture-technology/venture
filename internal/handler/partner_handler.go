package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/usecase/partner"
)

type PartnerHandler struct {
	PartnerUseCase *partner.PartnerUseCase
}

func NewPartnerHandler(pu *partner.PartnerUseCase) *PartnerHandler {
	return &PartnerHandler{
		PartnerUseCase: pu,
	}
}

func (ph *PartnerHandler) Get(c *gin.Context) {

	id := c.Param("id")

	partner, err := ph.PartnerUseCase.Get(c, &id)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "partner not found"})
		return
	}

	c.JSON(http.StatusOK, partner)

}

func (ph *PartnerHandler) FindAllByCnh(c *gin.Context) {

	cnh := c.Param("cnh")

	partners, err := ph.PartnerUseCase.FindAllByCnh(c, &cnh)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "partner not found"})
		return
	}

	c.JSON(http.StatusOK, partners)

}

func (ph *PartnerHandler) FindAllByCnpj(c *gin.Context) {

	cnpj := c.Param("cnpj")

	partners, err := ph.PartnerUseCase.FindAllByCnpj(c, &cnpj)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "partner not found"})
		return
	}

	c.JSON(http.StatusOK, partners)

}

func (ph *PartnerHandler) Delete(c *gin.Context) {

	id := c.Param("id")

	err := ph.PartnerUseCase.Delete(c, &id)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "partner not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted w successfully"})

}
