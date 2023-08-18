package ggMapService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	GoogleMapsAPIKey = "AIzaSyCf39yvpQbCspF2e0fHvxolCPT4FT4fJiA"
)

type GeoResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

type ggMapService struct {
}

type IGGService interface {
	GetCoordinateByGG(placeName string) (latitude, longitude float64, err error)
}

func NewGGMapService() IGGService {
	return &ggMapService{}
}

func (g *ggMapService) GetCoordinateByGG(placeName string) (latitude, longitude float64, err error) {
	endpoint := "https://maps.googleapis.com/maps/api/geocode/json"
	values := url.Values{}
	values.Add("address", placeName)
	values.Add("key", GoogleMapsAPIKey)
	url := fmt.Sprintf("%s?%s", endpoint, values.Encode())

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error when getting coordinates: %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error when reading coordinates: %v", err)
	}

	var geoResponse GeoResponse
	err = json.Unmarshal(body, &geoResponse)
	if err != nil {
		log.Fatalf("Error when binding json: %v", err)
	}

	if len(geoResponse.Results) > 0 {
		lat := geoResponse.Results[0].Geometry.Location.Lat
		lng := geoResponse.Results[0].Geometry.Location.Lng
		fmt.Printf("Latitude: %f, Longitude: %f\n", lat, lng)
	} else {
		fmt.Println("No results found")
	}
	return geoResponse.Results[0].Geometry.Location.Lat, geoResponse.Results[0].Geometry.Location.Lng, nil
}
