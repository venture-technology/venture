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
)

type KidController struct {
}

func NewKidController() *KidController {
	return &KidController{}
}

func (ch *KidController) PostV1CreateKid(c *gin.Context) {
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

	var requestParams entity.Kid
	if err := c.BindJSON(&requestParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		infra.App.Logger.Infof(fmt.Sprintf("error: %v", err.Error()))
		return
	}

	requestParams.ResponsibleCPF = cpf

	usecase := usecase.NewCreateKidUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.CreateKid(&requestParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar novo filho"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ch *KidController) GetV1GetKid(c *gin.Context) {
	rg := c.Param("rg")

	usecase := usecase.NewGetKidUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	kid, err := usecase.GetKid(&rg)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, kid)
}

func (ch *KidController) GetV1ListKids(c *gin.Context) {
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

	usecase := usecase.NewListKidsUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	kids, err := usecase.ListKids(&cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, kids)
}

func (ch *KidController) PatchV1UpdateController(c *gin.Context) {
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

	rg := c.Param("rg")
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewUpdateKidUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateKid(rg, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno ao atualizar informações"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (ch *KidController) DeleteV1DeleteKid(c *gin.Context) {
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

	rg := c.Param("rg")
	usecase := usecase.NewDeleteKidUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeleteKid(&rg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar filho"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
