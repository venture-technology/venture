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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data entity.DistanceMatrixResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	kmFloat, err := parseDistanceText(data)
	if err != nil {
		return nil, err
	}

	return &kmFloat, nil
}

func parseDistanceText(data entity.DistanceMatrixResponse) (float64, error) {
	if len(data.Rows) == 0 || len(data.Rows[0].Elements) == 0 {
		return 0, fmt.Errorf("invalid distance data")
	}

	distance := data.Rows[0].Elements[0].Distance.Text
	distance = strings.TrimSpace(strings.Replace(distance, "km", "", 1))
	distance = strings.ReplaceAll(distance, ",", ".")

	return strconv.ParseFloat(distance, 64)
}
