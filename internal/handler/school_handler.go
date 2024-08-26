package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/school"
	"github.com/venture-technology/venture/pkg/utils"
)

type SchoolHandler struct {
	schoolUseCase *school.SchoolUseCase
}

func NewSchoolHandler(schoolUseCase *school.SchoolUseCase) *SchoolHandler {
	return &SchoolHandler{
		schoolUseCase: schoolUseCase,
	}
}

func (sh *SchoolHandler) Create(c *gin.Context) {

	var input entity.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	validatecnpj := input.ValidateCnpj()
	if !validatecnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cnpj is invalid"})
	}

	input.Password = utils.HashPassword(input.Password)

	err := sh.schoolUseCase.Create(c, &input)

	if err != nil {
		log.Printf("error to create school: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating school"})
		return
	}

	log.Print("school create was successful")

	c.JSON(http.StatusCreated, input)

}

func (sh *SchoolHandler) Get(c *gin.Context) {

	cnpj := c.Param("cnpj")

	school, err := sh.schoolUseCase.Get(c, &cnpj)

	if err != nil {
		log.Printf("error while found school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school don't found"})
		return
	}

	c.JSON(http.StatusOK, school)

}

func (sh *SchoolHandler) FindAll(c *gin.Context) {

	schools, err := sh.schoolUseCase.FindAll(c)

	if err != nil {
		log.Printf("error while found schools: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "schools don't found"})
		return
	}

	c.JSON(http.StatusOK, schools)

}

func (sh *SchoolHandler) Update(c *gin.Context) {

	cnpj := c.Param("cnpj")

	var input entity.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	input.CNPJ = cnpj

	err := sh.schoolUseCase.Update(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error at update"})
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (sh *SchoolHandler) Delete(c *gin.Context) {

	cnpj := c.Param("cnpj")

	err := sh.schoolUseCase.Delete(c, &cnpj)

	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cnpj)

	c.JSON(http.StatusOK, gin.H{"message": "school deleted w successfully"})

}
