package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type InviteController struct {
}

func NewInviteController() *InviteController {
	return &InviteController{}
}

// @Summary Envia um convite
// @Description Cria um convite entre escola e motorista
// @Tags Invites
// @Accept json
// @Produce json
// @Param invite body entity.Invite true "Dados do convite"
// @Success 201 {object} nil
// @Failure 400 {object} map[string]string
// @Router /invite [post]
func (ih *InviteController) PostV1SendInvite(httpContext *gin.Context) {
	var requestParams entity.Invite
	if err := httpContext.BindJSON(&requestParams); err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewSendInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.SendInvite(&requestParams)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	httpContext.JSON(http.StatusCreated, http.NoBody)
}

// @Summary Lista convites do motorista
// @Description Retorna todos os convites enviados para o motorista
// @Tags Invites
// @Produce json
// @Param cnh path string true "CNH do motorista"
// @Success 200 {array} []value.DriverListInvite
// @Failure 400 {object} map[string]string
// @Router /driver/invites/{cnh} [get]
func (ih *InviteController) GetV1DriverListInvite(httpContext *gin.Context) {
	cnh := httpContext.Param("cnh")

	usecase := usecase.NewListDriverInvitesUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	invites, err := usecase.ListDriverInvites(cnh)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	httpContext.JSON(http.StatusOK, invites)
}

// @Summary Lista convites da escola
// @Description Retorna todos os convites enviados pela escola
// @Tags Invites
// @Produce json
// @Param cnpj path string true "CNPJ da escola"
// @Success 200 {array} []value.SchoolListInvite
// @Failure 400 {object} map[string]string
// @Router /school/invites/{cnpj} [get]
func (ih *InviteController) GetV1SchoolListInvite(httpContext *gin.Context) {
	cnpj := httpContext.Param("cnpj")

	usecase := usecase.NewListSchoolInvitesUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	invites, err := usecase.ListSchoolInvites(cnpj)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	httpContext.JSON(http.StatusOK, invites)
}

// @Summary Aceita convite
// @Description Aceita um convite existente
// @Tags Invites
// @Produce json
// @Param id path string true "ID do convite"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string
// @Router /invite/{id}/accept [patch]
func (ih *InviteController) PatchV1AcceptInvite(httpContext *gin.Context) {
	id := httpContext.Param("id")

	usecase := usecase.NewAcceptInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.AcceptInvite(id)

	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar aceitar convite"})
		return
	}

	httpContext.JSON(http.StatusOK, http.NoBody)
}

// @Summary Recusa convite
// @Description Recusa e deleta um convite existente
// @Tags Invites
// @Produce json
// @Param id path string true "ID do convite"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Router /invite/{id}/decline [delete]
func (ih *InviteController) DeleteV1DeclineInvite(httpContext *gin.Context) {
	id := httpContext.Param("id")

	usecase := usecase.NewDeclineInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.DeclineInvite(id)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar deletar convite"})
		return
	}

	httpContext.JSON(http.StatusNoContent, http.NoBody)
}
