package addresses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/entity"
)

type GoogleAdapter struct {
}

func NewGoogleAdapter() *GoogleAdapter {
	return &GoogleAdapter{}
}

func (ga *GoogleAdapter) GetDistance(origin, destination string) (*float64, error) {
	endpoint := viper.GetString("GOOGLE_MATRIX_DISTANCE_API_URL")

	params := url.Values{
		"units":        {"metric"},
		"origins":      {origin},
		"destinations": {destination},
		"key":          {viper.GetString("GOOGLE_CLOUD_SECRET_KEY")},
	}

	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data entity.DistanceMatrixResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	distance := data.Rows[0].Elements[0].Distance.Text
	distance = strings.TrimSpace(strings.Replace(distance, "km", "", 1))
	kmFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return nil, err
	}

	return &kmFloat, err
}
