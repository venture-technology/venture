package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

type IGoogleAdapter interface {
	GetDistance(origin, destination string) (*float64, error)
}

type GoogleAdapter struct {
}

func NewGoogleAdapter() *GoogleAdapter {
	return &GoogleAdapter{}
}

func (ga *GoogleAdapter) GetDistance(origin, destination string) (*float64, error) {
	conf := config.Get()

	endpoint := conf.GoogleCloudSecret.EndpointMatrixDistance

	params := url.Values{
		"units":        {"metric"},
		"origins":      {origin},
		"destinations": {destination},
		"key":          {conf.GoogleCloudSecret.ApiKey},
	}

	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	var data entity.DistanceMatrixResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Print(err.Error())
		return nil, err
	}

	if data.Status != "OK" {
		log.Print("Erro na API:", data.Status)
	}

	distance := data.Rows[0].Elements[0].Distance.Text

	distance = strings.TrimSpace(strings.Replace(distance, "km", "", 1))

	kmFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return nil, err
	}

	return &kmFloat, err
}
