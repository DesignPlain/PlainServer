package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"golang.org/x/exp/slices"
)

type APIController interface {
	AddCORSHeader(c *gin.Context)
	DeployStack(c *gin.Context)
	DestroyStack(c *gin.Context)
	SaveState(c *gin.Context)
	GetState(c *gin.Context)
}

type APIControllerImp struct {
	DataStore    []DeploymentResource
	CurrentStack []string
	Count        int64
}

func (_controllerStatus *APIControllerImp) AddCORSHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
}

func (_controllerStatus *APIControllerImp) DeployStack(c *gin.Context) {
	_controllerStatus.Count += 1
	if _controllerStatus.DataStore != nil {
		for _, data := range _controllerStatus.DataStore {
			switch data.ResourceType {
			case Simple_Storage_Service:
				if !slices.Contains(_controllerStatus.CurrentStack, data.Name+"_ST") {
					err := PulumiStackUp(data.Name + "_ST")
					if err != nil {
						_controllerStatus.CurrentStack = append(_controllerStatus.CurrentStack, data.Name+"_ST")
						Log(DEBUG, "Adding new data to stack "+strings.Join(_controllerStatus.CurrentStack, " "))
					}
				} else {
					Log(DEBUG, "Stack already contains the data")
				}
			}
		}

		c.JSON(403, "_controllerStatus.DataStore")
		return
	}

	Log(DEBUG, "Deploying stack")

	c.JSON(200, gin.H{"message": "Bucket Address"})
}

func (_controllerStatus *APIControllerImp) DestroyStack(c *gin.Context) {
	Log(DEBUG, "Destroying stack")
	for i := 0; i < len(_controllerStatus.CurrentStack); i++ {
		Log(DEBUG, _controllerStatus.CurrentStack)
		PulumiStackDestroy(_controllerStatus.CurrentStack[i])
		_controllerStatus.CurrentStack = append(_controllerStatus.CurrentStack[:i], _controllerStatus.CurrentStack[i+1:]...)
	}
	c.JSON(200, "_controllerStatus.DataStore")
}

func (_controllerStatus *APIControllerImp) SaveState(c *gin.Context) {
	var depResource []DeploymentResource

	if err := c.BindJSON(&depResource); err != nil {
		fmt.Println(err)
		return
	}

	Log(DEBUG, "Saving state")
	_controllerStatus.Count += 1

	_controllerStatus.DataStore = depResource

	data, _ := json.Marshal(_controllerStatus.DataStore)
	os.WriteFile("./.localStore/Database", data, 0644)

	//Log(DEBUG, depResource)
	c.JSON(200, gin.H{"message": "Done"})
}

func (_controllerStatus *APIControllerImp) GetState(c *gin.Context) {
	Log(DEBUG, "Returning Stored state")
	_controllerStatus.Count += 1
	if _controllerStatus.DataStore != nil {
		c.JSON(200, _controllerStatus.DataStore)
		return
	} else {
		data, _ := os.ReadFile("./.localStore/Database")
		json.Unmarshal(data, &_controllerStatus.DataStore)
		Log(DEBUG, _controllerStatus.DataStore)
	}

	c.JSON(200, gin.H{"Error": "State Not Found"})
}

func PulumiStackUp(resName string) error {
	ctx := context.Background()
	workDir := filepath.Join("/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/gcp")
	data, err := os.ReadFile("/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/main-form-398518-44df76f5c489.json")
	if err != nil {
		Log(ERROR, err)
		return err
	}
	envvars := auto.EnvVars(map[string]string{
		"PULUMI_CONFIG_PASSPHRASE": "test",
		"GOOGLE_CREDENTIALS":       string(data[:]),
	})

	cfg := auto.ConfigMap{
		"gcp:project": {Value: "main-form-398518"},
	}

	stackName := auto.FullyQualifiedStackName("organization", "Stack", resName)
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, workDir, envvars)
	stack.SetAllConfig(ctx, cfg)

	if err != nil {
		Log(ERROR, err)
		return err
	}

	// aa, err := stack.Info(ctx)
	// if err != nil {
	// 	Log(ERROR, err)
	// 	return
	// }

	// Log(DEBUG, aa)

	// previewResult, err := stack.Preview(ctx)
	// if err != nil {
	// 	Log(ERROR, err)
	// 	return
	// }

	// Log(DEBUG, previewResult)

	stackUpResult, err := stack.Up(ctx)
	Log(DEBUG, "Creating Stack "+stackName)
	if err != nil {
		Log(ERROR, err)
		return err
	}
	Log(DEBUG, stackUpResult)

	return nil
}

func PulumiStackDestroy(resName string) {
	ctx := context.Background()
	workDir := filepath.Join("/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/gcp")
	data, err := os.ReadFile("/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/main-form-398518-44df76f5c489.json")
	if err != nil {
		Log(ERROR, err)
		return
	}
	envvars := auto.EnvVars(map[string]string{
		"PULUMI_CONFIG_PASSPHRASE": "test",
		"GOOGLE_CREDENTIALS":       string(data[:]),
	})

	cfg := auto.ConfigMap{
		"gcp:project": {Value: "main-form-398518"},
	}

	stackName := auto.FullyQualifiedStackName("organization", "Stack", resName)
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, workDir, envvars)
	stack.SetAllConfig(ctx, cfg)

	if err != nil {
		Log(ERROR, err)
		return
	}

	stackDestroyResult, err := stack.Destroy(ctx)
	Log(DEBUG, "Destroying Stack "+stackName)
	if err != nil {
		Log(ERROR, err)
	}
	_ = stack.Workspace().RemoveStack(ctx, stackName)
	Log(DEBUG, stackDestroyResult)
}
