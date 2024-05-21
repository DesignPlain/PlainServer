package main

import (
	"DesignSphere_Server/src/model"
	. "DesignSphere_Server/src/model"
	utils "DesignSphere_Server/src/utils"
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

var backendGlobalContext = make(map[string]any)

type ProjectData struct {
	ProjectName         string
	GCP_APIKeyFileName  string
	AWS_AccessKeyId     string
	AWS_SecretAccessKey string
}

type APIController struct {
	DataStore     []DeploymentResource
	CurrentStack  []string
	ProjectConfig map[string]ProjectData
	Count         int64
}

func (_controllerStatus *APIController) ValidateState() bool {
	return len(_controllerStatus.ProjectConfig) > 0
}

func (_controllerStatus *APIController) AddCORSHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
}

func (_controllerStatus *APIController) InitServer() {
	utils.Log(utils.INFO, "Initialializing the server.")
	data, err := os.ReadFile("./.localStore/ServerState")
	if err == nil {
		json.Unmarshal(data, &backendGlobalContext)
		_controllerStatus.ProjectConfig["GCP"] = ProjectData{
			ProjectName:         backendGlobalContext["GCP_ProjectName"].(string),
			GCP_APIKeyFileName:  backendGlobalContext["GCP_APIFileName"].(string),
			AWS_AccessKeyId:     backendGlobalContext["AWS_AccessKeyId"].(string),
			AWS_SecretAccessKey: backendGlobalContext["AWS_AccessKey"].(string),
		}
		utils.Log(utils.DEBUG, _controllerStatus.ProjectConfig)
	}
}

func SaveServeState() {
	data, _ := json.Marshal(backendGlobalContext)
	os.WriteFile("./.localStore/ServerState", data, 0644)
}

func (_controllerStatus *APIController) DeployStack(c *gin.Context) {
	// if !_controllerStatus.ValidateState() {
	// 	c.String(http.StatusForbidden, "Please configure the cloud project")
	// 	return
	// }

	_controllerStatus.Count += 1
	var resourceOutputs []ResourceOutput

	utils.Log(utils.DEBUG, _controllerStatus.DataStore)

	// TODO: update the logic and datastructure to respect the resource dependency
	//       eg: VPCs should be created before the subnet and compute instances
	//       will have to add a DAG.
	if _controllerStatus.DataStore != nil {
		doDeploy := true
		var nextResourceIndex int
		pendingIndexQueue := getNextResourceIndexToDeploy(_controllerStatus.DataStore)
		utils.Log(utils.DEBUG, "AWS VPC deployment")
		var serarray []int
		for i := pendingIndexQueue.Front(); i != nil; i = i.Next() {
			serarray = append(serarray, i.Value.(int))
		}

		backendGlobalContext["stackOrder"] = serarray

		SaveServeState()
		for doDeploy {
			if pendingIndexQueue.Len() > 0 {
				front := pendingIndexQueue.Front()
				nextResourceIndex = front.Value.(int)
				pendingIndexQueue.Remove(front)
			} else {
				doDeploy = false
			}
			data := &_controllerStatus.DataStore[nextResourceIndex]

			switch data.ProviderType {
			case GCP:
				{
					projectConfig := _controllerStatus.ProjectConfig["GCP"]
					var resourceConfigModel ResourceModel
					if !slices.Contains(_controllerStatus.CurrentStack, data.Id.String()) {
						resourceConfigModel = model.CreateResourceModel(projectConfig.ProjectName, "", data.ResourceType, data.Name, data.ResourceConfig)

					} else {
						utils.Log(utils.DEBUG, "Stack already contains the data")
					}

					if resourceConfigModel.Name != "" {
						serializedResourceConfigYaml, err := yaml.Marshal(resourceConfigModel)
						utils.Log(utils.DEBUG, string(serializedResourceConfigYaml))
						if err != nil {
							log.Fatal(err)
						}

						resourceOutputs, err = PulumiStackUp(GCP, data.Id.String(), string(serializedResourceConfigYaml), &projectConfig)

						if err == nil {
							_controllerStatus.CurrentStack = append(_controllerStatus.CurrentStack, data.Id.String())
							backendGlobalContext["CurrentStack"] = _controllerStatus.CurrentStack
							SaveServeState()
							utils.Log(utils.DEBUG, "Adding new data to stack "+strings.Join(_controllerStatus.CurrentStack, " "))
						} else {
							utils.Log(utils.DEBUG, "Stack Up failed with error: "+err.Error())
						}
						data.Outputs = resourceOutputs
					}

					break
				}
			case AWS:
				{
					utils.Log(utils.DEBUG, "AWS VPC deployment switch 1")
					projectConfig := _controllerStatus.ProjectConfig["AWS"]
					utils.Log(utils.DEBUG, projectConfig)
					var resourceConfigModel ResourceModel

					if !slices.Contains(_controllerStatus.CurrentStack, data.Id.String()) {
						resourceConfigModel = model.CreateResourceModel(projectConfig.ProjectName, "", data.ResourceType, data.Name, data.ResourceConfig)
					} else {
						utils.Log(utils.DEBUG, "Stack already contains the data")
					}

					if resourceConfigModel.Resources != nil {
						serializedResourceConfigYaml, err := yaml.Marshal(resourceConfigModel)
						utils.Log(utils.DEBUG, string(serializedResourceConfigYaml))
						if err != nil {
							log.Fatal(err)
						}

						resourceOutputs, err = PulumiStackUp(AWS, data.Id.String(), string(serializedResourceConfigYaml), &projectConfig)

						if err == nil {
							_controllerStatus.CurrentStack = append(_controllerStatus.CurrentStack, data.Id.String())
							backendGlobalContext["CurrentStack"] = _controllerStatus.CurrentStack
							SaveServeState()
							utils.Log(utils.DEBUG, "Adding new data to stack "+strings.Join(_controllerStatus.CurrentStack, " "))
						} else {
							utils.Log(utils.DEBUG, "Stack Up failed with error: ")
							utils.Log(utils.DEBUG, err)
						}
						data.Outputs = resourceOutputs
					}

					break
				}
			// case AZURE:
			// 	{
			// 	}
			default:
				c.String(http.StatusForbidden, "Please provide a valid Cloud provider")
			}
		}
	}

	utils.Log(utils.DEBUG, _controllerStatus.DataStore)
	utils.Log(utils.DEBUG, "Deploying stack")
	c.JSON(200, _controllerStatus.DataStore)
}

func getNextResourceIndexToDeploy(dataStore []DeploymentResource) *list.List {
	var first *list.Element
	var pendingIndexQueue *list.List = list.New()

	for i := 0; i < len(dataStore); i++ {
		if dataStore[i].ResourceType == Virtual_Private_Cloud {
			if first == nil {
				first = pendingIndexQueue.PushFront(i)
			} else {
				pendingIndexQueue.PushFront(i)
			}
		} else if dataStore[i].ResourceType == Subnet {
			if first == nil {
				pendingIndexQueue.PushBack(i)
			} else {
				pendingIndexQueue.InsertBefore(i, first)
			}
		} else {
			pendingIndexQueue.PushBack(i)
		}
	}
	utils.Log(utils.DEBUG, "Deployment Plan is:")
	fmt.Println(pendingIndexQueue)
	return pendingIndexQueue
}

func (_controllerStatus *APIController) UploadProjectConfig(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")

	projectName := c.Request.Form["projectName"][0]
	log.Println(url.PathEscape(projectName))
	filename := url.PathEscape(strings.Replace(projectName, ".", "", -1)) + ".json"
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, ROOTDIR+filename)

	_controllerStatus.ProjectConfig["GCP"] = ProjectData{
		ProjectName:        projectName,
		GCP_APIKeyFileName: filename,
	}

	backendGlobalContext["GCP_ProjectName"] = projectName
	backendGlobalContext["GCP_APIFileName"] = filename

	SaveServeState()
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (_controllerStatus *APIController) DestroyStack(c *gin.Context) {
	if !_controllerStatus.ValidateState() {
		c.String(http.StatusForbidden, "Please configure the cloud project")
		return
	}

	projectConfig := _controllerStatus.ProjectConfig["GCP"]
	utils.Log(utils.DEBUG, "Destroying stack")
	utils.Log(utils.DEBUG, _controllerStatus.CurrentStack)
	doDestroy := true
	pendingIndexQueuear := backendGlobalContext["stackOrder"].([]int)
	nextResourceIndex := 0
	var pendingIndexQueue *list.List = list.New()

	_controllerStatus.CurrentStack = backendGlobalContext["CurrentStack"].([]string)

	for i := 0; i < len(pendingIndexQueuear); i++ {
		pendingIndexQueue.PushBack(i)
	}

	for doDestroy {
		if pendingIndexQueue.Len() > 0 {
			back := pendingIndexQueue.Back()
			nextResourceIndex = back.Value.(int)
			utils.Log(utils.DEBUG, "Destroy index: "+fmt.Sprint(nextResourceIndex))
			pendingIndexQueue.Remove(back)
		} else {
			doDestroy = false
		}

		utils.Log(utils.DEBUG, _controllerStatus.CurrentStack)
		PulumiStackDestroy(_controllerStatus.CurrentStack[nextResourceIndex], &projectConfig)
	}

	// TODO: Check if any of the stack destroy logic failed
	_controllerStatus.CurrentStack = nil

	c.JSON(200, "_controllerStatus.DataStore")
}

func (_controllerStatus *APIController) SaveState(c *gin.Context) {
	var depResource []DeploymentResource

	var jsonData []byte
	if c.Request.Body != nil {
		jsonData, _ = io.ReadAll(c.Request.Body)
	}

	if err := json.Unmarshal(jsonData, &depResource); err != nil {
		fmt.Println(err)
		return
	}

	utils.Log(utils.DEBUG, "Saving state")
	_controllerStatus.Count += 1

	_controllerStatus.DataStore = depResource

	data, _ := json.Marshal(_controllerStatus.DataStore)
	os.WriteFile("./.localStore/Database", data, 0644)

	utils.Log(utils.DEBUG, string(data))
	c.JSON(200, gin.H{"message": "Done"})
}

func (_controllerStatus *APIController) GetState(c *gin.Context) {
	utils.Log(utils.DEBUG, "Returning Stored state")
	_controllerStatus.Count += 1
	if _controllerStatus.DataStore == nil {
		data, err := os.ReadFile("./.localStore/Database")
		if err != nil {
			_controllerStatus.DataStore = []DeploymentResource{}
		} else {
			json.Unmarshal(data, &_controllerStatus.DataStore)
		}
	}

	utils.Log(utils.DEBUG, _controllerStatus.DataStore)
	c.JSON(200, _controllerStatus.DataStore)
}

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func PulumiStackUp(provider ProviderType, stackId string, yamlConfigModel string, projectConfig *ProjectData) ([]ResourceOutput, error) {
	ctx := context.Background()
	// TODO: Make directory for different providers
	var providerDir string
	if provider == GCP {
		providerDir = "gcp/"
	} else if provider == AWS {
		providerDir = "aws/"
	}

	workDir := filepath.Join(ROOTDIR + providerDir + stackId)
	data, err := os.ReadFile(ROOTDIR + projectConfig.GCP_APIKeyFileName)

	if err != nil {
		utils.Log(utils.ERROR, err)
		return nil, err
	}

	envvars := auto.EnvVars(map[string]string{
		"PULUMI_CONFIG_PASSPHRASE": "test",
		"GOOGLE_CREDENTIALS":       string(data[:]),
		"AWS_ACCESS_KEY_ID":        projectConfig.AWS_AccessKeyId,
		"AWS_SECRET_ACCESS_KEY":    projectConfig.AWS_SecretAccessKey,
	})

	// TODO: handle case for different providers
	cfg := auto.ConfigMap{
		"gcp:project": {Value: projectConfig.ProjectName},
		"aws:region":  {Value: "us-west-2"},
	}

	utils.Log(utils.DEBUG, workDir)

	// TODO: if value updated, rewrite the yaml
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

			os.WriteFile(workDir+"/Pulumi.yaml", []byte(yamlConfigModel), 0644)
		} else {
			return nil, err
		}
	}

	stackName := auto.FullyQualifiedStackName("organization", projectConfig.ProjectName, stackId)
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, workDir, envvars)
	stack.SetAllConfig(ctx, cfg)

	if err != nil {
		utils.Log(utils.ERROR, err)
		return nil, err
	}

	// aa, err := stack.Info(ctx)
	// if err != nil {
	// 	utils.Log(utils.ERROR, err)
	// 	return
	// }

	// utils.Log(utils.DEBUG, aa)

	// previewResult, err := stack.Preview(ctx)
	// if err != nil {
	// 	utils.Log(utils.ERROR, err)
	// 	return
	// }

	// utils.Log(utils.DEBUG, previewResult)

	stackUpResult, err := stack.Up(ctx)
	utils.Log(utils.DEBUG, "Creating Stack "+stackName)
	if err != nil {
		utils.Log(utils.ERROR, err)
		utils.Log(utils.DEBUG, stackUpResult)
		return nil, err
	}
	utils.Log(utils.DEBUG, stackUpResult)
	var resOuts []ResourceOutput
	for key, element := range stackUpResult.Outputs {
		//fmt.Println("Key:", key, "=>", "Element:", element.Value)
		a := ResourceOutput{
			Name:  key,
			Value: element.Value,
		}
		resOuts = append(resOuts, a)
	}

	utils.Log(utils.DEBUG, resOuts)
	return resOuts, nil
}

func PulumiStackDestroy(stackId string, projectConfig *ProjectData) {
	ctx := context.Background()
	workDir := filepath.Join(ROOTDIR + "gcp/" + stackId)
	data, err := os.ReadFile(ROOTDIR + projectConfig.GCP_APIKeyFileName)
	if err != nil {
		utils.Log(utils.ERROR, err)
		return
	}
	envvars := auto.EnvVars(map[string]string{
		"PULUMI_CONFIG_PASSPHRASE": "test",
		"GOOGLE_CREDENTIALS":       string(data[:]),
	})

	cfg := auto.ConfigMap{
		"gcp:project": {Value: projectConfig.ProjectName},
	}

	stackName := auto.FullyQualifiedStackName("organization", projectConfig.ProjectName, stackId)
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, workDir, envvars)
	stack.SetAllConfig(ctx, cfg)

	if err != nil {
		utils.Log(utils.ERROR, err)
		return
	}

	stackDestroyResult, err := stack.Destroy(ctx)
	utils.Log(utils.DEBUG, "Destroying Stack "+stackName)
	if err != nil {
		utils.Log(utils.ERROR, err)
	}
	_ = stack.Workspace().RemoveStack(ctx, stackName)
	utils.Log(utils.DEBUG, stackDestroyResult)
}
