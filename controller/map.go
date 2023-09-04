package controller

import (
	"VApplicationBE/plugin-service/ggMapService"
	"VApplicationBE/plugin-service/mapBoxService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mapHandler struct {
	mapService mapBoxService.IMapService
}

func NewHandlerMap() mapHandler {
	return mapHandler{}
}

type reqMapArgs struct {
	CustomerID string `json:"customerId"`
	RoadName   string `json:"roadName"`
}

type Geo struct {
	CustomerID string `json:"customerId" gorm:"primaryKey"`
	RoadName   string `json:"roadName"`
	Latitude   float64
	Longitude  float64
}

func (m *mapHandler) GetCoordinatesByMapBox(c *gin.Context) {
	var request reqMapArgs
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can not get the request")
		return
	}
	getMapService := mapBoxService.NewMapService()
	latitude, longitude, err := getMapService.GetCoordinatesByMapBox(request.RoadName)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error in getting coords")
		return
	}
	geo := Geo{
		CustomerID: request.CustomerID,
		RoadName:   request.RoadName,
		Latitude:   latitude,
		Longitude:  longitude,
	}
	c.JSON(http.StatusCreated, geo)
}

func (m *mapHandler) GetCoordinatesByGG(c *gin.Context) {
	var request reqMapArgs
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can not get the request")
		return
	}
	getGGMapService := ggMapService.NewGGMapService()
	latitude, longitude, err := getGGMapService.GetCoordinateByGG(request.RoadName)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error in getting coords")
		return
	}
	geo := Geo{
		CustomerID: request.CustomerID,
		RoadName:   request.RoadName,
		Latitude:   latitude,
		Longitude:  longitude,
	}
	c.JSON(http.StatusCreated, geo)
}
