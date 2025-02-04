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

func (ih *InviteController) PostV1SendInvite(c *gin.Context) {
	var requestParams entity.Invite
	if err := c.BindJSON(&requestParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewSendInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.SendInvite(&requestParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ih *InviteController) GetV1DriverListInvite(c *gin.Context) {
	cnh := c.Param("cnh")

	usecase := usecase.NewListDriverInvitesUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	invites, err := usecase.ListDriverInvites(cnh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	c.JSON(http.StatusOK, invites)
}

func (ih *InviteController) GetV1SchoolListInvite(c *gin.Context) {
	cnpj := c.Param("cnpj")

	usecase := usecase.NewListSchoolInvitesUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	invites, err := usecase.ListSchoolInvites(cnpj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	c.JSON(http.StatusOK, invites)
}

func (ih *InviteController) PatchV1AcceptInvite(c *gin.Context) {
	id := c.Param("id")

	usecase := usecase.NewAcceptInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.AcceptInvite(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar aceitar convite"})
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (ih *InviteController) DeleteV1DeclineInvite(c *gin.Context) {
	id := c.Param("id")

	usecase := usecase.NewDeclineInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.DeclineInvite(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar deletar convite"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
