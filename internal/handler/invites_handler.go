package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/invite"
)

type InviteHandler struct {
	inviteUseCase *invite.InviteUseCase
}

func NewInviteHandler(iu *invite.InviteUseCase) *InviteHandler {
	return &InviteHandler{
		inviteUseCase: iu,
	}
}

func (ih *InviteHandler) Create(c *gin.Context) {

	var input entity.Invite

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ih.inviteUseCase.Create(c, &input)

	if err != nil {
		log.Printf("error while creating invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, &input)

}

func (ih *InviteHandler) Get(c *gin.Context) {

	inviteId := c.Param("id")

	id, err := uuid.Parse(inviteId)

	if err != nil {
		log.Printf("parse uuid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invite don't found"})
		return
	}

	invite, err := ih.inviteUseCase.Get(c, id)

	if err != nil {
		log.Printf("error while found invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invite don't found"})
		return
	}

	c.JSON(http.StatusOK, invite)

}

func (ih *InviteHandler) FindAllByCnh(c *gin.Context) {

	cnh := c.Param("cnh")

	invites, err := ih.inviteUseCase.FindAllByCnh(c, &cnh)

	if err != nil {
		log.Printf("invites don't found: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusAccepted, invites)

}

func (ih *InviteHandler) FindAllByCnpj(c *gin.Context) {

	cnpj := c.Param("cnpj")

	invites, err := ih.inviteUseCase.FindAllByCnpj(c, &cnpj)

	if err != nil {
		log.Printf("invites don't found: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusAccepted, invites)

}

func (ih *InviteHandler) Accept(c *gin.Context) {

	inviteId := c.Param("id")

	id, err := uuid.Parse(inviteId)

	if err != nil {
		log.Printf("parse uuid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invite don't found"})
		return
	}

	err = ih.inviteUseCase.Accept(c, id)

	if err != nil {
		log.Printf("error while accepting invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at accepting invite"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "invite accepted"})

}

func (ih *InviteHandler) Decline(c *gin.Context) {

	inviteId := c.Param("id")

	id, err := uuid.Parse(inviteId)

	if err != nil {
		log.Printf("parse uuid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invite don't found"})
		return
	}

	err = ih.inviteUseCase.Decline(c, id)

	if err != nil {
		log.Printf("error while deleting invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at deleting invite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invite declined w/ successfully"})

}
