package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		infra.App.Config,
	)

	requestParams.Password = utils.MakeHash(requestParams.Password)

	err := usecase.CreateResponsible(&requestParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar criar responsável"))
		return
	}

	c.JSON(http.StatusCreated, requestParams)
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

	usecase := usecase.NewUpdateResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.UpdateResponsible(cpf, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar atualizar as informações do responsável na stripe"))
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (rh *ResponsibleController) DeleteV1DeleteResponsbile(c *gin.Context) {
	cpf := c.Param("cpf")

	usecase := usecase.NewDeleteResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Config,
	)

	// buscando customerid do responsible
	err := usecase.DeleteResponsible(cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "ao tentar buscar a chave do cliente no stripe"))
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusNoContent, http.NoBody)
}
