package mapBoxService

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
)

type mapService struct {
}

type IMapService interface {
	GetCoordinatesByMapBox(placeName string) (latitude, longitude float64, err error)
}

func NewMapService() IMapService {
	return &mapService{}
}

type MapboxResponse struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

const MAPBOX_API_TOKEN = "pk.eyJ1IjoiaGFpZGFpZGFvIiwiYSI6ImNsbDB1ZzZxbjBrenMzZ28xYzlocWRoaWsifQ.RM7HKuCavQfV1p-Bo4V2mw"

func (m *mapService) GetCoordinatesByMapBox(placeName string) (latitude, longitude float64, err error) {
	apiURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s",
		url.QueryEscape(placeName), MAPBOX_API_TOKEN)

	status, body, err := fasthttp.Get(nil, apiURL)
	if err != nil {
		return 0, 0, err
	}

	if status != fasthttp.StatusOK {
		return 0, 0, fmt.Errorf("failed to get coordinates, got status code %d", status)
	}

	var response MapboxResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, 0, err
	}

	if len(response.Features) == 0 {
		return 0, 0, fmt.Errorf("no features found for place name: %s", placeName)
	}

	coords := response.Features[0].Geometry.Coordinates
	return coords[1], coords[0], nil
}
