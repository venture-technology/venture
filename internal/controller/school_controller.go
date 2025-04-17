package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/domain/service/middleware"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
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
	var requestParams entity.School
	if err := c.BindJSON(&requestParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	validatecnpj := requestParams.ValidateCnpj()
	if !validatecnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cnpj é inválido"})
		return
	}

	hash, err := utils.MakeHash(requestParams.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	requestParams.Password = hash

	usecase := usecase.NewCreateSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.CreateSchool(&requestParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar criar escola"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
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
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	middleware := middleware.NewSchoolMiddleware()

	middlewareResponse, err := middleware.GetSchoolFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.School.CNPJ != cnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewUpdateSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateSchool(cnpj, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno de servidor ao atualizar as informações da escola"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (sh *SchoolController) DeleteV1DeleteSchool(c *gin.Context) {
	cnpj := c.Param("cnpj")

	middleware := middleware.NewSchoolMiddleware()

	middlewareResponse, err := middleware.GetSchoolFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.School.CNPJ != cnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewDeleteSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeleteSchool(cnpj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar escola"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusNoContent, http.NoBody)
}

func (sh *SchoolController) PostV1LoginSchool(httpContext *gin.Context) {
	var requestParams entity.School
	if err := httpContext.BindJSON(&requestParams); err != nil {
		infra.App.Logger.Errorf(fmt.Sprintf("error on bind json: %v", err))
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	err := requestParams.ValidateLogin()
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewSchoolLoginUsecase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	token, err := usecase.LoginSchool(requestParams.Email, requestParams.Password)
	if err != nil {
		if err.Error() == "user not found" {
			httpContext.JSON(http.StatusUnauthorized, exceptions.InvalidBodyContentResponseError(err))
			return
		}
		httpContext.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao realizar login"))
		return
	}

	httpContext.SetCookie("token", token, 3600*24*30, "/", httpContext.Request.Host, false, true)
	httpContext.JSON(http.StatusOK, gin.H{"token": token})
}
