package main

import (
	"VApplicationBE/plugin-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// listen on port 8080

	handler := controller.NewHandlerMap()

	r.POST("/geo", handler.GetCoordinatesByMapBox)
	r.POST("/geo-with-gg", handler.GetCoordinatesByGG)
	r.Run(":8080")

}
