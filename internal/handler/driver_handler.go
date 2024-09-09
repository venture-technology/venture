package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/usecase/driver"
	"github.com/venture-technology/venture/pkg/utils"
)

type DriverHandler struct {
	driverUseCase *driver.DriverUseCase
}

func NewDriverHandler(du *driver.DriverUseCase) *DriverHandler {
	return &DriverHandler{
		driverUseCase: du,
	}
}

func (dh *DriverHandler) Create(c *gin.Context) {

	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.Password = utils.HashPassword(input.Password)

	err := dh.driverUseCase.Create(c, &input)
	if err != nil {
		log.Printf("error to create driver: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured qwhen creating driver"))
		return
	}

	log.Print("driver create was successful")

	c.JSON(http.StatusCreated, input)

}

func (dh *DriverHandler) Get(c *gin.Context) {
	cnh := c.Param("cnh")

	driver, err := dh.driverUseCase.Get(c, &cnh)
	if err != nil {
		log.Printf("error while found driver: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "driver not found"))
		return
	}

	c.JSON(http.StatusOK, driver)
}

func (dh *DriverHandler) Update(c *gin.Context) {

	cnh := c.Param("cnh")

	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CNH = cnh

	err := dh.driverUseCase.Update(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "internal server error at update"))
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (dh *DriverHandler) Delete(c *gin.Context) {

	cnh := c.Param("cnh")

	err := dh.driverUseCase.Delete(c, &cnh)
	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cnh)

	c.JSON(http.StatusOK, gin.H{"message": "driver deleted w successfully"})

}

func (dh *DriverHandler) SavePix(c *gin.Context) {

	cnh := c.Param("cnh")

	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CNH = cnh

	err := dh.driverUseCase.SavePix(c, &input)
	if err != nil {
		log.Printf("save pix error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to save pix"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "pix key created"})

}

func (dh *DriverHandler) SaveBank(c *gin.Context) {

	cnh := c.Param("cnh")

	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CNH = cnh

	err := dh.driverUseCase.SaveBank(c, &input)
	if err != nil {
		log.Printf("save pix error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to save pix"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bank infos created"})

}

func (dh *DriverHandler) GetGallery(c *gin.Context) {

	cnh := c.Param("cnh")

	links, err := dh.driverUseCase.GetGallery(c, &cnh)
	if err != nil {
		log.Printf("error to get gallery: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to get gallery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"images": links})

}
