package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/child"
)

type ChildHandler struct {
	childUseCase *child.ChildUseCase
}

func NewChildHandler(childUseCase *child.ChildUseCase) *ChildHandler {
	return &ChildHandler{
		childUseCase: childUseCase,
	}
}

func (ch *ChildHandler) Create(c *gin.Context) {

	var input entity.Child

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ch.childUseCase.Create(c, &input)

	if err != nil {
		log.Printf("error to create child: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating child"})
		return
	}

	log.Print("child created was successful")

	c.JSON(http.StatusCreated, input)
}

func (ch *ChildHandler) Get(c *gin.Context) {

	rg := c.Param("rg")

	child, err := ch.childUseCase.Get(c, &rg)

	if err != nil {
		log.Printf("error while found child: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "child don't found"})
		return
	}

	c.JSON(http.StatusOK, child)

}

func (ch *ChildHandler) FindAll(c *gin.Context) {

	cpf := c.Param("cpf")

	children, err := ch.childUseCase.FindAll(c, &cpf)

	if err != nil {
		log.Printf("error while found children: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "children don't found"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ch.childUseCase.Update(c, &input)

	if err != nil {
		log.Printf("update error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error at update"})
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ch *ChildHandler) Delete(c *gin.Context) {

	rg := c.Param("rg")

	err := ch.childUseCase.Delete(c, &rg)
	if err != nil {
		log.Printf("delete child error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted child"})
		return
	}

	log.Printf("deleted your account --> %v", rg)

	c.JSON(http.StatusOK, gin.H{"message": "child deleted w successfully"})

}
