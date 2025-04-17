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

func (rh *ResponsibleController) PostV1CreateResponsible(c *gin.Context) {
	var requestParams entity.Responsible
	if err := c.BindJSON(&requestParams); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewCreateResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	hash, err := utils.MakeHash(requestParams.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	requestParams.Password = hash

	err = usecase.CreateResponsible(&requestParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar criar responsável"))
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (rh *ResponsibleController) GetV1GetResponsible(c *gin.Context) {
	cpf := c.Param("cpf")

	usecase := usecase.NewGetResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	responsible, err := usecase.GetResponsible(cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsavel não encontrado"))
		return
	}

	c.JSON(http.StatusOK, responsible)
}

func (rh *ResponsibleController) PatchV1UpdateResponsible(c *gin.Context) {
	cpf := c.Param("cpf")
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	middleware := middleware.NewResponsibleMiddleware()

	middlewareResponse, err := middleware.GetResponsibleFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Responsible.CPF != cpf {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewUpdateResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateResponsible(cpf, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar atualizar as informações do responsável na stripe"))
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (rh *ResponsibleController) DeleteV1DeleteResponsbile(c *gin.Context) {
	cpf := c.Param("cpf")

	middleware := middleware.NewResponsibleMiddleware()

	middlewareResponse, err := middleware.GetResponsibleFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Responsible.CPF != cpf {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewDeleteResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	// buscando customerid do responsible
	err = usecase.DeleteResponsible(cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "ao tentar buscar a chave do cliente no stripe"))
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusNoContent, http.NoBody)
}

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
