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
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type SchoolController struct {
}

func NewSchoolController() *SchoolController {
	return &SchoolController{}
}

// @Summary Cria uma nova escola
// @Description Cria uma nova escola com os dados fornecidos
// @Tags Schools
// @Accept json
// @Produce json
// @Param school body entity.School true "Dados da escola"
// @Success 201 {object} value.GetSchool
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /school [post]
func (sh *SchoolController) PostV1CreateSchool(httpContext *gin.Context) {
	var requestParams entity.School
	if err := httpContext.BindJSON(&requestParams); err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	validatecnpj := requestParams.ValidateCnpj()
	if !validatecnpj {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "cnpj é inválido"})
		return
	}

	hash, err := utils.MakeHash(requestParams.Password)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	requestParams.Password = hash

	usecase := usecase.NewCreateSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.CreateSchool(&requestParams)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar criar escola"})
		return
	}

	httpContext.JSON(http.StatusCreated, value.MapSchoolEntityToResponse(requestParams))
}

// @Summary Busca escola
// @Description Retorna a escola buscada pelo seu documento principal
// @Tags Schools
// @Produce json
// @Param cnpj path string true "CNPJ da escola"
// @Success 200 {object} value.GetSchool
// @Failure 400 {object} map[string]string
// @Router /school/{cnpj} [get]
func (sh *SchoolController) GetV1GetSchool(httpContext *gin.Context) {
	cnpj := httpContext.Param("cnpj")

	usecase := usecase.NewGetSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	school, err := usecase.GetSchool(cnpj)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "escola não encontrada"})
		return
	}

	httpContext.JSON(http.StatusOK, school)
}

// @Summary Lista todas as escolas
// @Description Retorna uma lista de todas as escolas cadastradas
// @Tags Schools
// @Produce json
// @Success 200 {array} value.GetSchool
// @Failure 400 {object} map[string]string
// @Router /schools [get]
func (sh *SchoolController) GetV1ListSchool(httpContext *gin.Context) {
	usecase := usecase.NewListSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	schools, err := usecase.ListSchool()
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "nenhuma escola encontrada"})
		return
	}

	httpContext.JSON(http.StatusOK, schools)
}

// @Summary Atualiza uma escola
// @Description Atualiza os dados de uma escola pelo CNPJ
// @Tags Schools
// @Accept json
// @Produce json
// @Param cnpj path string true "CNPJ da escola"
// @Param data body map[string]interface{} true "Dados a serem atualizados"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /school/{cnpj} [patch]
func (sh *SchoolController) PatchV1UpdateSchool(httpContext *gin.Context) {
	cnpj := httpContext.Param("cnpj")
	var data map[string]interface{}
	if err := httpContext.BindJSON(&data); err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	middleware := middleware.NewSchoolMiddleware(
		infra.App.Config,
	)

	middlewareResponse, err := middleware.GetSchoolFromMiddleware(httpContext)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.School.CNPJ != cnpj {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewUpdateSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateSchool(cnpj, data)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"message": "erro interno de servidor ao atualizar as informações da escola"})
		return
	}

	httpContext.JSON(http.StatusNoContent, http.NoBody)
}

// @Summary Deleta uma escola
// @Description Deleta uma escola pelo CNPJ
// @Tags Schools
// @Produce json
// @Param cnpj path string true "CNPJ da escola"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /school/{cnpj} [delete]
func (sh *SchoolController) DeleteV1DeleteSchool(httpContext *gin.Context) {
	cnpj := httpContext.Param("cnpj")

	middleware := middleware.NewSchoolMiddleware(
		infra.App.Config,
	)

	middlewareResponse, err := middleware.GetSchoolFromMiddleware(httpContext)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.School.CNPJ != cnpj {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewDeleteSchoolUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeleteSchool(cnpj)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar escola"})
		return
	}

	httpContext.SetCookie("token", "", -1, "/", httpContext.Request.Host, false, true)
	httpContext.JSON(http.StatusNoContent, http.NoBody)
}

// @Summary Login de escola
// @Description Realiza o login de uma escola com email e senha
// @Tags Schools
// @Accept json
// @Produce json
// @Param credentials body entity.School true "Credenciais da escola (email e senha)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /school/login [post]
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
		infra.App.Config,
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
