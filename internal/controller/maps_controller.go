package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/domain/adapter"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type MapsController struct {
}

func NewMapsController() *MapsController {
	return &MapsController{}
}

func (mh *MapsController) CalculatePrice(c *gin.Context) {
	var input entity.MapPrice
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecase := usecase.NewCalculatePriceUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		adapter.NewGoogleAdapter(),
	)

	origin := fmt.Sprintf(
		"%s,%s,%s",
		input.Origin.Street,
		input.Origin.Number,
		input.Origin.ZIP,
	)

	destination := fmt.Sprintf(
		"%s,%s,%s",
		input.Destination.Street,
		input.Destination.Number,
		input.Destination.ZIP,
	)

	value, err := usecase.CalculatePrice(origin, destination, input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valor": value})
}
