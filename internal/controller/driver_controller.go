package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/domain/service/middleware"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
	"github.com/venture-technology/venture/pkg/utils"
)

type DriverController struct {
}

func NewDriverController() *DriverController {
	return &DriverController{}
}

func (dh *DriverController) PostV1Create(c *gin.Context) {
	var requestParams entity.Driver
	if err := c.BindJSON(&requestParams); err != nil {
		infra.App.Logger.Errorf(fmt.Sprintf("error on bind json: %v", err))
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	hash, err := utils.MakeHash(requestParams.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	requestParams.Password = hash

	usecase := usecase.NewCreateDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Bucket,
	)

	err = usecase.CreateDriver(&requestParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao realziar a criação do qrcode"))
		return
	}

	c.JSON(http.StatusCreated, http.NoBody)
}

func (dh *DriverController) GetV1GetDriver(c *gin.Context) {
	cnh := c.Param("cnh")

	usecase := usecase.NewGetDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Bucket,
	)

	driver, err := usecase.GetDriver(cnh)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "motorista não encontrado"))
		return
	}

	c.JSON(http.StatusOK, driver)
}

func (dh *DriverController) PatchV1UpdateDriver(c *gin.Context) {
	cnh := c.Param("cnh")
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	middleware := middleware.NewDriverMiddleware()

	middlewareResponse, err := middleware.GetDriverFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Driver.CNH != cnh {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewUpdateDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateDriver(cnh, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao realizar atualização das informações do motorista"))
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (dh *DriverController) DeleteV1DeleteDriver(c *gin.Context) {
	cnh := c.Param("cnh")

	middleware := middleware.NewDriverMiddleware()

	middlewareResponse, err := middleware.GetDriverFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Driver.CNH != cnh {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewDeleteDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeleteDriver(cnh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar motorista"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusNoContent, http.NoBody)
}

func (dh *DriverController) PostV1LoginDriver(httpContext *gin.Context) {
	var requestParams entity.Driver
	if err := httpContext.BindJSON(&requestParams); err != nil {
		infra.App.Logger.Errorf(fmt.Sprintf("error on bind json: %v", err))
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	err := requestParams.ValidateLogin()
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewDriverLoginUsecase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	token, err := usecase.LoginDriver(requestParams.Email, requestParams.Password)
	if err != nil {
		if err.Error() == "user not found" {
			httpContext.JSON(http.StatusUnauthorized, exceptions.InvalidBodyContentResponseError(err))
			return
		}
		httpContext.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao realizar login"))
		return
	}

	httpContext.SetCookie("token", token, 3600*24*30, "/", httpContext.Request.Host, false, true)
	httpContext.JSON(http.StatusOK, gin.H{"token": token})
}
