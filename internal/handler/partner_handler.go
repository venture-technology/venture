package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/usecase/partner"
	"go.uber.org/zap"
)

type PartnerHandler struct {
	PartnerUseCase *partner.PartnerUseCase
	logger         *zap.Logger
}

func NewPartnerHandler(pu *partner.PartnerUseCase, logger *zap.Logger) *PartnerHandler {
	return &PartnerHandler{
		PartnerUseCase: pu,
		logger:         logger,
	}
}

func (ph *PartnerHandler) Get(c *gin.Context) {

	id := c.Param("id")

	partner, err := ph.PartnerUseCase.Get(c, &id)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, partner)

}

func (ph *PartnerHandler) FindAllByCnh(c *gin.Context) {

	cnh := c.Param("cnh")

	partners, err := ph.PartnerUseCase.FindAllByCnh(c, &cnh)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, partners)

}

func (ph *PartnerHandler) FindAllByCnpj(c *gin.Context) {

	cnpj := c.Param("cnpj")

	partners, err := ph.PartnerUseCase.FindAllByCnpj(c, &cnpj)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, partners)

}

func (ph *PartnerHandler) Delete(c *gin.Context) {

	id := c.Param("id")

	err := ph.PartnerUseCase.Delete(c, &id)

	if err != nil {
		log.Printf("error while found partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parceiro não encontrado"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)

}
