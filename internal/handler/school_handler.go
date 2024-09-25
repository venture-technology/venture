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
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	validatecnpj := input.ValidateCnpj()
	if !validatecnpj {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cnpj é inválido"})
		return
	}

	input.Password = utils.HashPassword(input.Password)

	err := sh.schoolUseCase.Create(c, &input)

	if err != nil {
		log.Printf("error to create school: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar criar escola"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "escola não encontrada"})
		return
	}

	c.JSON(http.StatusOK, school)

}

func (sh *SchoolHandler) FindAll(c *gin.Context) {

	schools, err := sh.schoolUseCase.FindAll(c)

	if err != nil {
		log.Printf("error while found schools: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "nenhuma escola encontrada"})
		return
	}

	c.JSON(http.StatusOK, schools)

}

func (sh *SchoolHandler) Update(c *gin.Context) {

	cnpj := c.Param("cnpj")

	var input entity.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	input.CNPJ = cnpj

	err := sh.schoolUseCase.Update(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "erro interno de servidor ao atualizar as informações da escola"})
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusNoContent, http.NoBody)

}

func (sh *SchoolHandler) Delete(c *gin.Context) {

	cnpj := c.Param("cnpj")

	err := sh.schoolUseCase.Delete(c, &cnpj)

	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar escola"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cnpj)

	c.JSON(http.StatusNoContent, http.NoBody)

}
