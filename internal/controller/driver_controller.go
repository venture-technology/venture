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

// @Summary Cria um novo motorista
// @Tags Drivers
// @Accept json
// @Produce json
// @Param body body entity.Driver true "Dados do motorista"
// @Success 201 {object} value.GetDriver
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /driver [post]
func (dh *DriverController) PostV1Create(httpContext *gin.Context) {
	var requestParams entity.Driver
	if err := httpContext.BindJSON(&requestParams); err != nil {
		infra.App.Logger.Errorf(fmt.Sprintf("error on bind json: %v", err))
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	ok, errors := utils.ValidatePassword(requestParams.Password)
	if !ok {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	hash, err := utils.MakeHash(requestParams.Password)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
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
		httpContext.JSON(http.StatusInternalServerError, err)
		return
	}

	httpContext.JSON(http.StatusCreated, http.NoBody)
}

// @Summary Busca motorista
// @Description Retorna os dados de um motorista pelo CNH
// @Tags Drivers
// @Produce json
// @Param cnh path string true "CNH do motorista"
// @Success 200 {object} value.GetDriver
// @Failure 400 {object} map[string]string
// @Router /driver/{cnh} [get]
func (dh *DriverController) GetV1GetDriver(httpContext *gin.Context) {
	cnh := httpContext.Param("cnh")

	usecase := usecase.NewGetDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Bucket,
	)

	driver, err := usecase.GetDriver(cnh)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "motorista não encontrado"))
		return
	}

	httpContext.JSON(http.StatusOK, driver)
}

// @Summary Atualiza dados do motorista
// @Tags Drivers
// @Accept json
// @Produce json
// @Param cnh path string true "CNH do motorista"
// @Param body body object true "Campos a serem atualizados"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Router /driver/{cnh} [patch]
func (dh *DriverController) PatchV1UpdateDriver(httpContext *gin.Context) {
	cnh := httpContext.Param("cnh")
	var data map[string]interface{}
	if err := httpContext.BindJSON(&data); err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	middleware := middleware.NewDriverMiddleware()

	middlewareResponse, err := middleware.GetDriverFromMiddleware(httpContext)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Driver.CNH != cnh {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewUpdateDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.UpdateDriver(cnh, data)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao realizar atualização das informações do motorista"))
		return
	}

	httpContext.JSON(http.StatusNoContent, http.NoBody)
}

// @Summary Deleta motorista
// @Tags Drivers
// @Produce json
// @Param cnh path string true "CNH do motorista"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Router /driver/{cnh} [delete]
func (dh *DriverController) DeleteV1DeleteDriver(httpContext *gin.Context) {
	cnh := httpContext.Param("cnh")

	middleware := middleware.NewDriverMiddleware()

	middlewareResponse, err := middleware.GetDriverFromMiddleware(httpContext)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar buscar o responsável do middleware"))
		return
	}

	if middlewareResponse.Driver.CNH != cnh {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "access denied"})
		return
	}

	usecase := usecase.NewDeleteDriverUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err = usecase.DeleteDriver(cnh)
	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar motorista"})
		return
	}

	httpContext.SetCookie("token", "", -1, "/", httpContext.Request.Host, false, true)
	httpContext.JSON(http.StatusNoContent, http.NoBody)
}

// @Summary Login de motorista
// @Tags Drivers
// @Accept json
// @Produce json
// @Param body body entity.Driver true "Credenciais de login (email e senha)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /driver/login [post]
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
