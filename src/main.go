package main

import (
	"github.com/gin-gonic/gin"
)

const ROOTDIR = "/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/"

func main() {
	r := gin.Default()

	var apiController APIController
	_apiControllerState := &APIControllerImp{
		DataStore:     nil,
		CurrentStack:  nil,
		Count:         0,
		ProjectConfig: make(map[string]ProjectData),
	}

	apiController = _apiControllerState

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Middleware
	r.Use(apiController.AddCORSHeader)

	// Route Handlers
	r.POST("/deploy", apiController.DeployStack)
	r.POST("/state", apiController.SaveState)
	r.GET("/state", apiController.GetState)
	r.GET("/stack", apiController.DestroyStack)
	// r.POST("/preview", DeploymentHandler)

	r.POST("/uploadProjectConfig", apiController.UploadProjectConfig)

	// listen and serve on 0.0.0.0:8080
	r.Run()

}
