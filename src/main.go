package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	_apiController := &APIControllerImp{
		DataStore:    nil,
		CurrentStack: nil,
		Count:        0,
	}

	// Middleware
	r.Use(_apiController.AddCORSHeader)

	// Route Handlers
	r.POST("/deploy", _apiController.DeployStack)
	r.POST("/state", _apiController.SaveState)
	r.GET("/state", _apiController.GetState)
	r.GET("/stack", _apiController.DestroyStack)
	// r.POST("/preview", DeploymentHandler)

	// listen and serve on 0.0.0.0:8080
	r.Run()

}
