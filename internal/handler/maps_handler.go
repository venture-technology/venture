package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/usecase/maps"
)

type MapsHandler struct {
	mapsUseCase *maps.MapsUseCase
}

func NewMapsHandler(mapsUseCase *maps.MapsUseCase) *MapsHandler {
	return &MapsHandler{
		mapsUseCase: mapsUseCase,
	}
}

func (mh *MapsHandler) CalculatePrice(c *gin.Context) {

	var input entity.MapPrice

	if err := c.BindJSON(&input); err != nil {
		log.Printf("bind json error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	origin := fmt.Sprintf("%s,%s,%s", input.Origin.Street, input.Origin.Number, input.Origin.ZIP)
	destination := fmt.Sprintf("%s,%s,%s", input.Destination.Street, input.Destination.Number, input.Destination.ZIP)

	value, err := mh.mapsUseCase.CalculatePrice(c, origin, destination, input.Amount)

	if err != nil {
		log.Printf("calc price error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": value})

}
