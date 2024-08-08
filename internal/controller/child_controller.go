package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/service"
	"github.com/venture-technology/venture/models"
)

type ChildController struct {
	childservice *service.ChildService
}

func NewChildController(childservice *service.ChildService) *ChildController {
	return &ChildController{
		childservice: childservice,
	}
}

func (ct *ChildController) RegisterRoutes(router *gin.Engine) {

	api := router.Group("api/v1/child")

	api.POST("/:cpf", ct.CreateChild)
	api.GET("/:rg", ct.GetChild)
	api.GET("/all/:cpf", ct.FindAllChildren)
	api.PATCH("/:rg", ct.UpdateChild)
	api.DELETE("/:rg", ct.DeleteChild)

}

func (ct *ChildController) CreateChild(c *gin.Context) {

	var input models.Child

	cpf := c.Param("cpf")

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	input.Responsible.CPF = cpf

	err := ct.childservice.CreateChild(c, &input)

	if err != nil {
		log.Printf("error to create child: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating child"})
		return
	}

	log.Print("child created was successful")

	c.JSON(http.StatusCreated, input)
}

func (ct *ChildController) GetChild(c *gin.Context) {

	rg := c.Param("rg")

	child, err := ct.childservice.GetChild(c, &rg)

	if err != nil {
		log.Printf("error while found child: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "child don't found"})
		return
	}

	c.JSON(http.StatusOK, child)

}

func (ct *ChildController) FindAllChildren(c *gin.Context) {

	cpf := c.Param("cpf")

	children, err := ct.childservice.FindAllChildren(c, &cpf)

	if err != nil {
		log.Printf("error while found children: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "children don't found"})
		return
	}

	c.JSON(http.StatusOK, children)

}

func (ct *ChildController) UpdateChild(c *gin.Context) {

	rg := c.Param("rg")

	var input models.Child

	input.RG = rg

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.childservice.UpdateChild(c, &input)

	if err != nil {
		log.Printf("update error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error at update"})
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ct *ChildController) DeleteChild(c *gin.Context) {

	rg := c.Param("rg")

	err := ct.childservice.DeleteChild(c, &rg)
	if err != nil {
		log.Printf("delete child error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted child"})
		return
	}

	log.Printf("deleted your account --> %v", rg)

	c.JSON(http.StatusOK, gin.H{"message": "child deleted w successfully"})

}
