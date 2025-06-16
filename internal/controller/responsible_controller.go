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

type ResponsibleController struct {
}

func NewResponsibleController() *ResponsibleController {
	return &ResponsibleController{}
}

// @Summary Cria um novo responsável
// @Description Cria um novo responsável com os dados fornecidos
// @Tags Responsibles
// @Accept json
// @Produce json
// @Param responsible body entity.Responsible true "Dados do responsável"
// @Success 201 {object} value.GetResponsible
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /responsible [post]
func (rh *ResponsibleController) PostV1CreateResponsible(httpContext *gin.Context) {
	var requestParams entity.Responsible
	if err := httpContext.BindJSON(&requestParams); err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewCreateResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	ok, errors := utils.ValidatePassword(requestParams.Password)
	if !ok {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	hash, err := utils.MakeHash(requestParams.Password)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	requestParams.Password = hash

	err = usecase.CreateResponsible(&requestParams)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar criar responsável"))
		return
	}

	httpContext.JSON(http.StatusCreated, http.NoBody)
}

// @Summary Busca responsável
// @Description Retorna o responsável buscado pelo seu CPF
// @Tags Responsibles
// @Produce json
// @Param cpf path string true "CPF do responsável"
// @Success 200 {object} value.GetResponsible
// @Failure 400 {object} map[string]string
// @Router /responsible/{cpf} [get]
func (rh *ResponsibleController) GetV1GetResponsible(httpContext *gin.Context) {
	cpf := httpContext.Param("cpf")

	usecase := usecase.NewGetResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	responsible, err := usecase.GetResponsible(cpf)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsavel não encontrado"))
		return
	}

	httpContext.JSON(http.StatusOK, responsible)
}

// @Summary Atualiza um responsável
// @Description Atualiza os dados de um responsável pelo CPF
// @Tags Responsibles
// @Accept json
// @Produce json
// @Param cpf path string true "CPF do responsável"
// @Param data body map[string]interface{} true "Dados a serem atualizados"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /responsible/{cpf} [patch]
func (rh *ResponsibleController) PatchV1UpdateResponsible(httpContext *gin.Context) {
	cpf := httpContext.Param("cpf")
	var data map[string]interface{}
	if err := httpContext.BindJSON(&data); err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	middleware := middleware.NewResponsibleMiddleware()

	middlewareResponse, err := middleware.GetResponsibleFromMiddleware(httpContext)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Responsible.CPF != cpf {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewUpdateResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateResponsible(cpf, data)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar atualizar as informações do responsável na stripe"))
		return
	}

	httpContext.JSON(http.StatusNoContent, http.NoBody)
}

// @Summary Deleta um responsável
// @Description Deleta um responsável pelo CPF
// @Tags Responsibles
// @Produce json
// @Param cpf path string true "CPF do responsável"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /responsible/{cpf} [delete]
func (rh *ResponsibleController) DeleteV1DeleteResponsbile(httpContext *gin.Context) {
	cpf := httpContext.Param("cpf")

	middleware := middleware.NewResponsibleMiddleware()

	middlewareResponse, err := middleware.GetResponsibleFromMiddleware(httpContext)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Responsible.CPF != cpf {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewDeleteResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeleteResponsible(cpf)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "ao tentar buscar a chave do cliente no stripe"))
		return
	}

	httpContext.SetCookie("token", "", -1, "/", httpContext.Request.Host, false, true)
	httpContext.JSON(http.StatusNoContent, http.NoBody)
}

// @Summary Login de responsável
// @Description Realiza o login de um responsável com email e senha
// @Tags Responsibles
// @Accept json
// @Produce json
// @Param credentials body entity.Responsible true "Credenciais do responsável (email e senha)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /responsible/login [post]
func (rh *ResponsibleController) PostV1LoginResponsible(httpContext *gin.Context) {
	var requestParams entity.Responsible
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

	usecase := usecase.NewResponsibleLoginUsecase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	token, err := usecase.LoginResponsible(requestParams.Email, requestParams.Password)
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
