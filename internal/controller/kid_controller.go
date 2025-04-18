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

// PostV1CreateKid godoc
// @Summary      Criar novo filho
// @Description  Cria um novo filho vinculado ao CPF do responsável
// @Tags         Kids
// @Accept       json
// @Produce      json
// @Param        cpf   path      string       true  "CPF do responsável"
// @Param        body  body      entity.Kid   true  "Dados do filho"
// @Success 201 {object} value.GetKid
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /kid/{cpf} [post]
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

// GetV1GetKid godoc
// @Summary      Buscar filho
// @Description  Busca um filho pelo RG
// @Tags         Kids
// @Accept       json
// @Produce      json
// @Param        rg   path      string       true  "RG do filho"
// @Success      200  {object}  value.GetKid
// @Failure      400  {object}  map[string]interface{}
// @Router       /kid/{rg} [get]
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

// GetV1ListKids godoc
// @Summary      Listar filhos
// @Description  Lista todos os filhos de um responsável
// @Tags         Kids
// @Accept       json
// @Produce      json
// @Param        cpf   path      string       true  "CPF do responsável"
// @Success 200 {array} []value.ListKid
// @Failure 400 {object} map[string]string
// @Router       /kids/{cpf} [get]
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

// PatchV1UpdateController godoc
// @Summary      Atualizar filho
// @Description  Atualiza os dados de um filho específico
// @Tags         Kids
// @Accept       json
// @Produce      json
// @Param        cpf   path      string                true  "CPF do responsável"
// @Param        rg    path      string                true  "RG do filho"
// @Param        body  body      map[string]interface{} true "Campos para atualizar"
// @Success      204   {object}  nil
// @Failure      400   {object}  map[string]interface{}
// @Router       /kid/{cpf}/{rg} [patch]
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

// DeleteV1DeleteKid godoc
// @Summary      Deletar filho
// @Description  Remove um filho vinculado ao CPF do responsável
// @Tags         Kids
// @Accept       json
// @Produce      json
// @Param        cpf   path      string  true  "CPF do responsável"
// @Param        rg    path      string  true  "RG do filho"
// @Success      204   {object}  nil
// @Failure      400   {object}  map[string]interface{}
// @Router       /kid/{cpf}/{rg} [delete]
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
