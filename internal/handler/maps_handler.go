package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/maps"
	"go.uber.org/zap"
)

type MapsHandler struct {
	mapsUseCase *maps.MapsUseCase
	logger      *zap.Logger
}

func NewMapsHandler(mapsUseCase *maps.MapsUseCase, logger *zap.Logger) *MapsHandler {
	return &MapsHandler{
		mapsUseCase: mapsUseCase,
		logger:      logger,
	}
}

func (mh *MapsHandler) CalculatePrice(c *gin.Context) {
	var input entity.MapPrice

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	origin := fmt.Sprintf("%s,%s,%s", input.Origin.Street, input.Origin.Number, input.Origin.ZIP)
	destination := fmt.Sprintf("%s,%s,%s", input.Destination.Street, input.Destination.Number, input.Destination.ZIP)

	value, err := mh.mapsUseCase.CalculatePrice(c, origin, destination, input.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valor": value})
}
