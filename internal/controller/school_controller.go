package controller

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/service"
	"github.com/venture-technology/venture/models"
)

type ClaimsSchool struct {
	CNPJ string `json:"cnpj"`
	jwt.StandardClaims
}

type SchoolController struct {
	schoolservice *service.SchoolService
}

func NewSchoolController(schoolservice *service.SchoolService) *SchoolController {
	return &SchoolController{
		schoolservice: schoolservice,
	}
}

func (ct *SchoolController) RegisterRoutes(router *gin.Engine) {

	api := router.Group("api/v1/school")

	api.POST("/", ct.CreateSchool)        // criar uma escola
	api.GET("/:cnpj", ct.ReadSchool)      // buscar uma escola em especifico
	api.GET("/", ct.ReadAllSchools)       // buscar todas as escolas
	api.PATCH("/:cnpj", ct.UpdateSchool)  // atualizar algum dado especifico
	api.DELETE("/:cnpj", ct.DeleteSchool) // deletar propria conta
}

func (ct *SchoolController) CreateSchool(c *gin.Context) {

	var input models.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	validatecnpj := input.ValidateCnpj()
	if !validatecnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cnpj is invalid"})
	}

	err := ct.schoolservice.CreateSchool(c, &input)

	if err != nil {
		log.Printf("error to create school: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating school"})
		return
	}

	log.Print("school create was successful")

	c.JSON(http.StatusCreated, input)

}

func (ct *SchoolController) ReadSchool(c *gin.Context) {

	cnpj := c.Param("cnpj")

	school, err := ct.schoolservice.ReadSchool(c, &cnpj)

	if err != nil {
		log.Printf("error while found school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"school": school})

}

func (ct *SchoolController) ReadAllSchools(c *gin.Context) {

	schools, err := ct.schoolservice.ReadAllSchools(c)

	if err != nil {
		log.Printf("error while found schools: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "schools don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"schools": schools})

}

func (ct *SchoolController) UpdateSchool(c *gin.Context) {

	cnpj := c.Param("cnpj")

	var input models.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	input.CNPJ = cnpj

	validatecnpj := input.ValidateCnpj()
	if !validatecnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cnpj is invalid"})
	}

	err := ct.schoolservice.UpdateSchool(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error at update"})
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ct *SchoolController) DeleteSchool(c *gin.Context) {

	cnpj := c.Param("cnpj")

	err := ct.schoolservice.DeleteSchool(c, &cnpj)

	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cnpj)

	c.JSON(http.StatusOK, gin.H{"message": "school deleted w successfully"})

}

func (ct *SchoolController) AuthSchool(c *gin.Context) {

}
