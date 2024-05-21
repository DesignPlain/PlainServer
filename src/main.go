package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

const ROOTDIR = "/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/"

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		signal.Reset(syscall.SIGINT)
	}()

	r := gin.Default()

	apiController := &APIController{
		DataStore:     nil,
		CurrentStack:  nil,
		Count:         0,
		ProjectConfig: make(map[string]ProjectData),
	}

	apiController.InitServer()

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
