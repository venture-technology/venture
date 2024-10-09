package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/child"
	"go.uber.org/zap"
)

type ChildHandler struct {
	childUseCase *child.ChildUseCase
	logger       *zap.Logger
}

func NewChildHandler(childUseCase *child.ChildUseCase, logger *zap.Logger) *ChildHandler {
	return &ChildHandler{
		childUseCase: childUseCase,
		logger:       logger,
	}
}

func (ch *ChildHandler) Create(c *gin.Context) {

	var input entity.Child

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	err := ch.childUseCase.Create(c, &input)

	if err != nil {
		log.Printf("error to create child: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar novo filho"})
		return
	}

	log.Print("child created was successful")

	c.JSON(http.StatusCreated, http.NoBody)
}

func (ch *ChildHandler) Get(c *gin.Context) {

	rg := c.Param("rg")

	child, err := ch.childUseCase.Get(c, &rg)

	if err != nil {
		log.Printf("error while found child: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, child)

}

func (ch *ChildHandler) FindAll(c *gin.Context) {

	cpf := c.Param("cpf")

	children, err := ch.childUseCase.FindAll(c, &cpf)

	if err != nil {
		log.Printf("error while found children: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "filho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, children)

}

func (ch *ChildHandler) Update(c *gin.Context) {

	rg := c.Param("rg")

	var input entity.Child

	input.RG = rg

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	err := ch.childUseCase.Update(c, &input)

	if err != nil {
		log.Printf("update error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno ao atualizar informações"})
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusNoContent, http.NoBody)

}

func (ch *ChildHandler) Delete(c *gin.Context) {

	rg := c.Param("rg")

	err := ch.childUseCase.Delete(c, &rg)
	if err != nil {
		log.Printf("delete child error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar filho"})
		return
	}

	log.Printf("deleted your account --> %v", rg)

	c.JSON(http.StatusNoContent, http.NoBody)

}
