package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var input entity.Invite
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewSendInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.SendInvite(&input)
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

	c.JSON(http.StatusAccepted, invites)
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

	c.JSON(http.StatusAccepted, invites)
}

func (ih *InviteController) PatchV1AcceptInvite(c *gin.Context) {
	inviteId := c.Param("id")
	uuid, err := uuid.Parse(inviteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "convite não encontrado"})
		return
	}

	usecase := usecase.NewAcceptInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.AcceptInvite(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar aceitar convite"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ih *InviteController) DeleteV1DeclineInvite(c *gin.Context) {
	inviteId := c.Param("id")

	uuid, err := uuid.Parse(inviteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "convite não encontrado"})
		return
	}

	usecase := usecase.NewDeclineInviteUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeclineInvite(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar deletar convite"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
