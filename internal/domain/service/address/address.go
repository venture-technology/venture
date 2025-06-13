package address

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/venture-technology/venture/internal/entity"
)

const (
	urlBase     = "https://maps.googleapis.com/maps"
	urlDistance = urlBase + "/api/distancematrix/json"
)

type Address interface {
	// Distance can be used to get distance between addresses.
	Distance(origin, destination string) (*float64, error)
}

type address struct {
	client string
}

func NewAddress(client string) *address {
	return &address{
		client: client,
	}
}

func (a *address) Distance(origin, destination string) (*float64, error) {

	params := url.Values{
		"units":        {"metric"},
		"origins":      {origin},
		"destinations": {destination},
		"key":          {a.client},
	}

	url := fmt.Sprintf("%s?%s", urlDistance, params.Encode())
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
