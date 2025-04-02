package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type PartnerController struct {
}

func NewPartnerController() *PartnerController {
	return &PartnerController{}
}

func (ph *PartnerController) GetV1DriverListPartners(c *gin.Context) {
	cnh := c.Param("cnh")

	usecase := usecase.NewDriverListPartnersUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	partners, err := usecase.DriverListPartners(cnh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, partners)
}

func (ph *PartnerController) GetV1SchoolListPartners(c *gin.Context) {
	cnpj := c.Param("cnpj")

	usecase := usecase.NewSchoolListPartnersUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	partners, err := usecase.SchoolListPartners(cnpj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, partners)
}

func (ph *PartnerController) DeleteV1DeletePartner(c *gin.Context) {
	id := c.Param("id")

	usecase := usecase.NewDeletePartnerUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.DeletePartner(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
