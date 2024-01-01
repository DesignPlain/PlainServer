package main

import (
	gcp "DesignSphere_Server/src/resource/GCP"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	UploadProjectConfig(c *gin.Context)
}

type APIControllerImp struct {
	DataStore     []DeploymentResource
	CurrentStack  []string
	ProjectConfig map[string]ProjectData
	Count         int64
}

func (_controllerStatus *APIControllerImp) ValidateState() bool {
	return len(_controllerStatus.ProjectConfig) > 0
}

type ProjectData struct {
	ProjectName    string
	APIKeyFileName string
}

func (_controllerStatus *APIControllerImp) AddCORSHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
}

func (_controllerStatus *APIControllerImp) DeployStack(c *gin.Context) {
	if !_controllerStatus.ValidateState() {
		c.String(http.StatusForbidden, fmt.Sprintf("Please configure the cloud project"))
	}
	_controllerStatus.Count += 1
	//var ma = make(map[uuid.UUID]([]ResourceOutput))
	var resOutputs []ResourceOutput
	var err error = nil

	if _controllerStatus.DataStore != nil {
		//	for _, data := range _controllerStatus.DataStore {
		for i := 0; i < len(_controllerStatus.DataStore); i++ {
			data := &_controllerStatus.DataStore[i]
			switch data.ResourceType {
			case Simple_Storage_Service:
				if !slices.Contains(_controllerStatus.CurrentStack, data.Name+"_ST") {
					yamlConfig := gcp.CreateBucketModel(
						data.Id.String(),
						data.Name,
						[]string{data.ResourceConfig["Members"].(string)},
						fmt.Sprintf(`"%s"`, data.ResourceConfig["Role"].(string)),
						data.ResourceConfig["Location"].(string))

					resOutputs, err = PulumiStackUp(data.Id.String(), data.Name+"_ST", yamlConfig, _controllerStatus.ProjectConfig["GCP"])
					if err == nil {
						_controllerStatus.CurrentStack = append(_controllerStatus.CurrentStack, data.Name+"_ST")
						Log(DEBUG, "Adding new data to stack "+strings.Join(_controllerStatus.CurrentStack, " "))
					} else {
						Log(DEBUG, "Stack Up failed")
						Log(DEBUG, err)
					}
					data.Outputs = resOutputs

				} else {
					Log(DEBUG, "Stack already contains the data")
				}
			}
		}
	}

	// for i := 0; i < len(_controllerStatus.DataStore); i++ {
	// 	_controllerStatus.DataStore[i].Outputs = ma[_controllerStatus.DataStore[i].Id]
	// 	Log(DEBUG, ma[_controllerStatus.DataStore[i].Id])
	// }

	Log(DEBUG, _controllerStatus.DataStore)
	Log(DEBUG, "Deploying stack")
	c.JSON(200, _controllerStatus.DataStore)
}

func (_controllerStatus *APIControllerImp) UploadProjectConfig(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")

	projectName := c.Request.Form["projectName"][0]
	log.Println(url.PathEscape(projectName))
	filename := url.PathEscape(strings.Replace(projectName, ".", "", -1)) + ".json"
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, ROOTDIR+filename)

	_controllerStatus.ProjectConfig["GCP"] = ProjectData{
		ProjectName:    projectName,
		APIKeyFileName: filename,
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (_controllerStatus *APIControllerImp) DestroyStack(c *gin.Context) {
	if !_controllerStatus.ValidateState() {
		c.String(http.StatusForbidden, fmt.Sprintf("Please configure the cloud project"))
	}
	Log(DEBUG, "Destroying stack")
	Log(DEBUG, _controllerStatus.CurrentStack)
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

	Log(DEBUG, depResource)
	c.JSON(200, gin.H{"message": "Done"})
}

func (_controllerStatus *APIControllerImp) GetState(c *gin.Context) {
	Log(DEBUG, "Returning Stored state")
	_controllerStatus.Count += 1
	if _controllerStatus.DataStore == nil {
		data, err := os.ReadFile("./.localStore/Database")
		if err != nil {
			_controllerStatus.DataStore = []DeploymentResource{}
		} else {
			json.Unmarshal(data, &_controllerStatus.DataStore)
		}
	}

	Log(DEBUG, _controllerStatus.DataStore)
	c.JSON(200, _controllerStatus.DataStore)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func PulumiStackUp(stackId string, resName string, yamlconfig string, projectConfig ProjectData) ([]ResourceOutput, error) {
	ctx := context.Background()
	workDir := filepath.Join(ROOTDIR + "gcp/" + stackId)
	data, err := os.ReadFile(ROOTDIR + projectConfig.APIKeyFileName)
	if err != nil {
		Log(ERROR, err)
		return nil, err
	}
	envvars := auto.EnvVars(map[string]string{
		"PULUMI_CONFIG_PASSPHRASE": "test",
		"GOOGLE_CREDENTIALS":       string(data[:]),
	})

	cfg := auto.ConfigMap{
		"gcp:project": {Value: projectConfig.ProjectName},
	}

	Log(DEBUG, workDir)

	_, err = os.Stat(workDir + "/Pulumi.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			// File or directory does not exist
			_, err = os.Stat(workDir)
			if err != nil {
				if os.IsNotExist(err) {
					if err := os.Mkdir(workDir, os.ModePerm); err != nil {
						panic(err)
					}
				}
			}

			f, err := os.Create(workDir + "/Pulumi.yaml")
			check(err)

			defer f.Close()

			_, err = f.WriteString(yamlconfig)
			check(err)

			f.Sync()
		} else {
			// Some other error. The file may or may not exist
		}
	}

	stackName := auto.FullyQualifiedStackName("organization", stackId, resName)
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, workDir, envvars)
	stack.SetAllConfig(ctx, cfg)

	if err != nil {
		Log(ERROR, err)
		return nil, err
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
		return nil, err
	}
	//Log(DEBUG, stackUpResult)
	var resOuts []ResourceOutput
	for key, element := range stackUpResult.Outputs {
		//fmt.Println("Key:", key, "=>", "Element:", element.Value)
		a := ResourceOutput{
			Name:  key,
			Value: element.Value,
		}
		resOuts = append(resOuts, a)
	}

	Log(DEBUG, resOuts)
	return resOuts, nil
}

func PulumiStackDestroy(resName string) {
	ctx := context.Background()
	workDir := filepath.Join(ROOTDIR + "gcp")
	data, err := os.ReadFile(ROOTDIR + "main-form-398518-44df76f5c489.json")
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
