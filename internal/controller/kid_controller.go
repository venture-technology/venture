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
func (ch *KidController) PostV1CreateKid(httpContext *gin.Context) {
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

	var requestParams entity.Kid
	if err := httpContext.BindJSON(&requestParams); err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
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
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar novo filho"})
		return
	}

	httpContext.JSON(http.StatusCreated, http.NoBody)
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
func (ch *KidController) GetV1GetKid(httpContext *gin.Context) {
	rg := httpContext.Param("rg")

	usecase := usecase.NewGetKidUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	kid, err := usecase.GetKid(&rg)

	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	httpContext.JSON(http.StatusOK, kid)
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
func (ch *KidController) GetV1ListKids(httpContext *gin.Context) {
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

	usecase := usecase.NewListKidsUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	kids, err := usecase.ListKids(&cpf)

	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	httpContext.JSON(http.StatusOK, kids)
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
func (ch *KidController) PatchV1UpdateController(httpContext *gin.Context) {
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

	rg := httpContext.Param("rg")
	var data map[string]interface{}
	if err := httpContext.BindJSON(&data); err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewUpdateKidUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateKid(rg, data)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"message": "erro interno ao atualizar informações"})
		return
	}

	httpContext.JSON(http.StatusNoContent, http.NoBody)
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
func (ch *KidController) DeleteV1DeleteKid(httpContext *gin.Context) {
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

	rg := httpContext.Param("rg")
	usecase := usecase.NewDeleteKidUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeleteKid(&rg)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar filho"})
		return
	}

	httpContext.JSON(http.StatusNoContent, http.NoBody)
}
