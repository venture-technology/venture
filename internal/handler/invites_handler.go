package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/invite"
	"go.uber.org/zap"
)

type InviteHandler struct {
	inviteUseCase *invite.InviteUseCase
	logger        *zap.Logger
}

func NewInviteHandler(iu *invite.InviteUseCase, logger *zap.Logger) *InviteHandler {
	return &InviteHandler{
		inviteUseCase: iu,
		logger:        logger,
	}
}

func (ih *InviteHandler) Create(c *gin.Context) {
	var input entity.Invite

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	err := ih.inviteUseCase.Create(c, &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ih *InviteHandler) Get(c *gin.Context) {
	inviteId := c.Param("id")

	id, err := uuid.Parse(inviteId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "convite não encontrado"})
		return
	}

	invite, err := ih.inviteUseCase.Get(c, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "convite não encontrado"})
		return
	}

	c.JSON(http.StatusOK, invite)
}

func (ih *InviteHandler) FindAllByCnh(c *gin.Context) {
	cnh := c.Param("cnh")

	invites, err := ih.inviteUseCase.FindAllByCnh(c, &cnh)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	c.JSON(http.StatusAccepted, invites)
}

func (ih *InviteHandler) FindAllByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	invites, err := ih.inviteUseCase.FindAllByCnpj(c, &cnpj)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno no servidor"})
		return
	}

	c.JSON(http.StatusAccepted, invites)
}

func (ih *InviteHandler) Accept(c *gin.Context) {
	inviteId := c.Param("id")

	id, err := uuid.Parse(inviteId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "convite não encontrado"})
		return
	}

	err = ih.inviteUseCase.Accept(c, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar aceitar convite"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ih *InviteHandler) Decline(c *gin.Context) {
	inviteId := c.Param("id")

	id, err := uuid.Parse(inviteId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "convite não encontrado"})
		return
	}

	err = ih.inviteUseCase.Decline(c, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro interno no servidor ao tentar deletar convite"})
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}
