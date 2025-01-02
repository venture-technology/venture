package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
	"github.com/venture-technology/venture/pkg/utils"
)

type SchoolController struct {
}

func NewSchoolController() *SchoolController {
	return &SchoolController{}
}

func (sh *SchoolController) PostV1CreateSchool(c *gin.Context) {
	var input entity.School
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	validatecnpj := input.ValidateCnpj()
	if !validatecnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cnpj é inválido"})
		return
	}

	input.Password = utils.MakeHash(input.Password)

	usecase := usecase.NewCreateSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.CreateSchool(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar criar escola"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (sh *SchoolController) GetV1GetSchool(c *gin.Context) {
	cnpj := c.Param("cnpj")

	usecase := usecase.NewGetSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	school, err := usecase.GetSchool(cnpj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "escola não encontrada"})
		return
	}

	c.JSON(http.StatusOK, school)
}

func (sh *SchoolController) GetV1ListSchool(c *gin.Context) {
	usecase := usecase.NewListSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	schools, err := usecase.ListSchool()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nenhuma escola encontrada"})
		return
	}

	c.JSON(http.StatusOK, schools)
}

func (sh *SchoolController) PatchV1UpdateSchool(c *gin.Context) {
	cnpj := c.Param("cnpj")

	var input entity.School
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	input.CNPJ = cnpj

	usecase := usecase.NewUpdateSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.UpdateSchool(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno de servidor ao atualizar as informações da escola"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (sh *SchoolController) DeleteV1DeleteSchool(c *gin.Context) {
	cnpj := c.Param("cnpj")

	usecase := usecase.NewDeleteSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.DeleteSchool(cnpj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar escola"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusNoContent, http.NoBody)
}
