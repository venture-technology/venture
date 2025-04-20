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

// @Summary Lista parceiros do motorista
// @Description Retorna todos os parceiros vinculados ao motorista
// @Tags Partners
// @Produce json
// @Param cnh path string true "CNH do motorista"
// @Success 200 {array} []value.DriverListPartners
// @Failure 400 {object} map[string]string
// @Router /driver/partners/{cnh} [get]
func (ph *PartnerController) GetV1DriverListPartners(httpContext *gin.Context) {
	cnh := httpContext.Param("cnh")

	usecase := usecase.NewDriverListPartnersUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	partners, err := usecase.DriverListPartners(cnh)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, partners)
}

// @Summary Lista parceiros da escola
// @Description Retorna todos os parceiros vinculados à escola
// @Tags Partners
// @Produce json
// @Param cnpj path string true "CNPJ da escola"
// @Success 200 {array} []value.SchoolListPartners
// @Failure 400 {object} map[string]string
// @Router /school/partners/{cnpj} [get]
func (ph *PartnerController) GetV1SchoolListPartners(httpContext *gin.Context) {
	cnpj := httpContext.Param("cnpj")

	usecase := usecase.NewSchoolListPartnersUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	partners, err := usecase.SchoolListPartners(cnpj)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, partners)
}

// @Summary Deleta parceria
// @Description Remove uma parceria entre motorista e escola
// @Tags Partners
// @Produce json
// @Param id path string true "ID da parceria"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Router /partner/{id} [delete]
func (ph *PartnerController) DeleteV1DeletePartner(httpContext *gin.Context) {
	id := httpContext.Param("id")

	usecase := usecase.NewDeletePartnerUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.DeletePartner(id)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusNoContent, http.NoBody)
}
