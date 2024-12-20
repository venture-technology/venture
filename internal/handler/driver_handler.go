package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/usecase/driver"
	"github.com/venture-technology/venture/pkg/utils"
	"go.uber.org/zap"
)

type DriverHandler struct {
	driverUseCase *driver.DriverUseCase
	logger        *zap.Logger
}

func NewDriverHandler(du *driver.DriverUseCase, logger *zap.Logger) *DriverHandler {
	return &DriverHandler{
		driverUseCase: du,
		logger:        logger,
	}
}

func (dh *DriverHandler) Create(c *gin.Context) {
	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.Password = utils.HashPassword(input.Password)

	err := dh.driverUseCase.Create(c, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao realziar a criação do qrcode"))
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (dh *DriverHandler) Get(c *gin.Context) {
	cnh := c.Param("cnh")

	driver, err := dh.driverUseCase.Get(c, &cnh)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "motorista não encontrado"))
		return
	}

	c.JSON(http.StatusOK, driver)
}

func (dh *DriverHandler) Update(c *gin.Context) {
	cnh := c.Param("cnh")

	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CNH = cnh

	err := dh.driverUseCase.Update(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao realizar atualização das informações do motorista"))
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (dh *DriverHandler) Delete(c *gin.Context) {
	cnh := c.Param("cnh")

	err := dh.driverUseCase.Delete(c, &cnh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar motorista"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusNoContent, http.NoBody)
}

func (dh *DriverHandler) SavePix(c *gin.Context) {
	cnh := c.Param("cnh")

	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CNH = cnh

	err := dh.driverUseCase.SavePix(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao salvar chave pix"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (dh *DriverHandler) SaveBank(c *gin.Context) {
	cnh := c.Param("cnh")

	var input entity.Driver

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CNH = cnh

	err := dh.driverUseCase.SaveBank(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao salvar informações da conta bancária"})
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (dh *DriverHandler) GetGallery(c *gin.Context) {
	cnh := c.Param("cnh")

	links, err := dh.driverUseCase.GetGallery(c, &cnh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao buscar galeria de imagens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"images": links})
}
