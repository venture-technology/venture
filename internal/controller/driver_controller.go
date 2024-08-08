package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/service"
	"github.com/venture-technology/venture/models"
)

type DriverController struct {
	driverservice *service.DriverService
}

func NewDriverController(driverservice *service.DriverService) *DriverController {
	return &DriverController{
		driverservice: driverservice,
	}
}

func (ct *DriverController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("api/v1/driver")

	api.POST("/", ct.CreateDriver)
	api.GET("/:cnh", ct.GetDriver)
	api.PATCH("/:cnh", ct.UpdateDriver)
	api.DELETE("/:cnh", ct.DeleteDriver)
}

func (ct *DriverController) CreateDriver(c *gin.Context) {
	var input models.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	urlQrCode, err := ct.driverservice.CreateAndSaveQrCode(c, input.CNH)
	if err != nil {
		log.Printf("error to create QrCode: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "an error occured qwhen creating QrCode"))
		return
	}

	input.QrCode = urlQrCode

	err = ct.driverservice.CreateDriver(c, &input)
	if err != nil {
		log.Printf("error to create driver: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured qwhen creating driver"))
		return
	}

	log.Print("driver create was successful")

	c.JSON(http.StatusCreated, input)
}

func (ct *DriverController) GetDriver(c *gin.Context) {
	cnh := c.Param("cnh")

	driver, err := ct.driverservice.GetDriver(c, &cnh)
	if err != nil {
		log.Printf("error while found driver: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "driver not found"))
		return
	}

	c.JSON(http.StatusOK, driver)
}

func (ct *DriverController) UpdateDriver(c *gin.Context) {

	cnh := c.Param("cnh")

	var input models.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CNH = cnh

	err := ct.driverservice.UpdateDriver(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "internal server error at update"))
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})
}

func (ct *DriverController) DeleteDriver(c *gin.Context) {

	cnh := c.Param("cpf")

	err := ct.driverservice.DeleteDriver(c, &cnh)
	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cnh)

	c.JSON(http.StatusOK, gin.H{"message": "driver deleted w successfully"})
}
