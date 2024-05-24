package server

import (
	. "DesignSphere_Server/src/server/model"
	"DesignSphere_Server/src/utils"
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

const ROOTDIR = "/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/"

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

func StartServer() {
	r := gin.Default()

	apiController := &APIController{
		DataStore:     nil,
		CurrentStack:  nil,
		Count:         0,
		ProjectConfig: make(map[string]ProjectData),
	}

	apiController.InitServer()

	r.MaxMultipartMemory = 8 << 20

	r.Use(apiController.AddCORSHeader)

	r.POST("/deploy", apiController.DeployStack)
	r.POST("/state", apiController.SaveState)
	r.GET("/state", apiController.GetState)
	r.GET("/stack", apiController.DestroyStack)

	r.POST("/uploadProjectConfig", apiController.UploadProjectConfig)
	r.GET("/getProjectConfig", apiController.GetProjectConfig)

	r.Run()
}

func (_controllerStatus *APIController) InitServer() {
	utils.Log(utils.INFO, "Initialializing the server.")
	data, err := os.ReadFile("./.localStore/ServerState")
	if err == nil {
		utils.Log(utils.INFO, "Syncing server state")
		json.Unmarshal(data, &backendGlobalContext)

		_controllerStatus.ProjectConfig["GCP"] = ProjectData{
			ProjectName:        backendGlobalContext["GCP_ProjectName"].(string),
			GCP_APIKeyFileName: backendGlobalContext["GCP_APIFileName"].(string),
		}

		_controllerStatus.ProjectConfig["AWS"] = ProjectData{
			AWS_AccessKeyId:     backendGlobalContext["AWS_AccessKeyId"].(string),
			AWS_SecretAccessKey: backendGlobalContext["AWS_SecretKey"].(string),
		}
		utils.Log(utils.DEBUG, _controllerStatus.ProjectConfig)
	}

	var depResource []DeploymentResource
	data, err = os.ReadFile("./.localStore/Database")
	if err == nil {
		utils.Log(utils.INFO, "Syncing resource state")
		json.Unmarshal(data, &depResource)
		_controllerStatus.DataStore = depResource
	}

}

func (_controllerStatus *APIController) UploadProjectConfig(c *gin.Context) {
	file, _ := c.FormFile("file")
	projectName := c.Request.Form.Get("projectName")

	awsKeyid := c.Request.Form.Get("awsKeyid")
	awsSecretKey := c.Request.Form.Get("awsSecretKey")

	if file != nil {
		filename := url.PathEscape(strings.Replace(projectName, ".", "", -1)) + ".json"
		if filename != "" {
			c.SaveUploadedFile(file, ROOTDIR+filename)
			backendGlobalContext["GCP_APIFileName"] = filename
		}
	}

	if projectName != "" {
		backendGlobalContext["GCP_ProjectName"] = projectName
	}

	_, c1 := backendGlobalContext["GCP_ProjectName"]
	_, c2 := backendGlobalContext["GCP_APIFileName"]

	if c1 || c2 {
		_controllerStatus.ProjectConfig["GCP"] = ProjectData{
			ProjectName:        backendGlobalContext["GCP_ProjectName"].(string),
			GCP_APIKeyFileName: backendGlobalContext["GCP_APIFileName"].(string),
		}
	}

	if awsKeyid != "" {
		backendGlobalContext["AWS_AccessKeyId"] = awsKeyid
	}
	if awsSecretKey != "" {
		backendGlobalContext["AWS_SecretKey"] = awsSecretKey
	}

	_, c1 = backendGlobalContext["AWS_AccessKeyId"]
	_, c2 = backendGlobalContext["AWS_SecretKey"]

	if c1 || c2 {
		_controllerStatus.ProjectConfig["AWS"] = ProjectData{
			ProjectName:        backendGlobalContext["AWS_AccessKeyId"].(string),
			GCP_APIKeyFileName: backendGlobalContext["AWS_SecretKey"].(string),
		}
	}

	SaveServeState()
	c.String(http.StatusOK, "Done")
}

func (_controllerStatus *APIController) GetProjectConfig(c *gin.Context) {
	var projectConfig = ProjectData{}

	if _, p := backendGlobalContext["AWS_AccessKeyId"]; p {
		projectConfig.AWS_AccessKeyId = backendGlobalContext["AWS_AccessKeyId"].(string)
	}

	if _, p := backendGlobalContext["AWS_SecretKey"]; p {
		//projectConfig.AWS_SecretAccessKey = backendGlobalContext["AWS_SecretKey"].(string)
		projectConfig.AWS_SecretAccessKey = "****************"

	}

	if _, p := backendGlobalContext["GCP_ProjectName"]; p {
		projectConfig.ProjectName = backendGlobalContext["GCP_ProjectName"].(string)
	}

	if _, p := backendGlobalContext["GCP_APIFileName"]; p {
		projectConfig.GCP_APIKeyFileName = backendGlobalContext["GCP_APIFileName"].(string)
	}

	c.JSON(http.StatusOK, projectConfig)
}

// Deploy and cloud resource stack
//
// TODO: Currently we deploy all the resources togther, we will still need this feature
// but will also require API to deploy the resources separatly.
// TODO: Fix deployment order related bugs
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
		var deploymentList []int

		pendingIndexQueue := getNextResourceIndexToDeploy(_controllerStatus.DataStore)
		for i := pendingIndexQueue.Front(); i != nil; i = i.Next() {
			deploymentList = append(deploymentList, i.Value.(int))
		}

		backendGlobalContext["stackOrder"] = deploymentList
		SaveServeState()

		for doDeploy {
			if pendingIndexQueue.Len() > 0 {
				front := pendingIndexQueue.Front()
				nextResourceIndex = front.Value.(int)
				pendingIndexQueue.Remove(front)
			} else {
				doDeploy = false
				break
			}
			data := &_controllerStatus.DataStore[nextResourceIndex]
			utils.Log(utils.DEBUG, "Processing resource: "+data.Id.String()+" "+data.Name)
			switch data.ProviderType {
			case GCP:
				{
					projectConfig := _controllerStatus.ProjectConfig["GCP"]
					var resourceConfigModel ResourceModel
					if !slices.Contains(_controllerStatus.CurrentStack, data.Id.String()) {
						resourceConfigModel = CreateResourceModel("DesignSphereStack", data.ResourceType.String(), data.ResourceType, data.Name, data.ResourceConfig)

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
							utils.Log(utils.ERROR, "Stack Up failed with error: "+err.Error())
							// c.JSON(403, err.Error())
							// return
						}
						data.Outputs = resourceOutputs
					}
				}
			case AWS:
				{
					projectConfig := _controllerStatus.ProjectConfig["AWS"]
					var resourceConfigModel ResourceModel
					if !slices.Contains(_controllerStatus.CurrentStack, data.Id.String()) {
						resourceConfigModel = CreateResourceModel("DesignSphereStack", data.ResourceType.String(), data.ResourceType, data.Name, data.ResourceConfig)
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
							utils.Log(utils.ERROR, "Stack Up failed with error: "+err.Error())
						}
						data.Outputs = resourceOutputs
					}
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

// API to bring down the deployed stack
//
// TODO: Will also need to create an API to destroy the resources separately
func (_controllerStatus *APIController) DestroyStack(c *gin.Context) {
	if !_controllerStatus.ValidateState() {
		c.String(http.StatusForbidden, "Please configure the cloud project")
		return
	}

	projectConfig := _controllerStatus.ProjectConfig["GCP"]
	utils.Log(utils.DEBUG, "Destroying stack")
	utils.Log(utils.DEBUG, _controllerStatus.CurrentStack)

	doDestroy := true
	nextResourceIndex := 0
	var pendingIndexQueue *list.List = list.New()

	pendingIndexQueuear, ok := backendGlobalContext["stackOrder"].([]int)

	if !ok {
		c.JSON(200, "No Stack to deploy")
		return
	}

	_controllerStatus.CurrentStack, ok = backendGlobalContext["CurrentStack"].([]string)
	if !ok {
		c.JSON(200, "Stack is empty")
		return

	}

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
			break
		}

		utils.Log(utils.DEBUG, _controllerStatus.CurrentStack)
		PulumiStackDestroy(_controllerStatus.CurrentStack[nextResourceIndex], &projectConfig)
	}

	// TODO: Check if any of the stack destroy logic failed
	_controllerStatus.CurrentStack = nil

	c.JSON(200, "_controllerStatus.DataStore")
}

// Save UI rendering state and config changes
func (_controllerStatus *APIController) SaveState(c *gin.Context) {
	var depResource []DeploymentResource
	var jsonData []byte

	if c.Request.Body == nil {
		c.JSON(401, gin.H{"message": "Empty body"})
	}

	jsonData, _ = io.ReadAll(c.Request.Body)
	if err := json.Unmarshal(jsonData, &depResource); err != nil {
		utils.Log(utils.ERROR, err)
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

// Get UI state with resource config
//
// TODO: Split this API to manage the UI rendering state and resource config changes separately
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

func PulumiStackUp(provider ProviderType, stackId string, yamlConfigModel string, projectConfig *ProjectData) ([]ResourceOutput, error) {
	ctx := context.Background()

	var envMap = map[string]string{"PULUMI_CONFIG_PASSPHRASE": "test"}

	var providerDir string
	if provider == GCP {
		providerDir = "gcp/"

		data, err := os.ReadFile(ROOTDIR + projectConfig.GCP_APIKeyFileName)
		if err != nil {
			utils.Log(utils.ERROR, err)
			return nil, err
		}

		envMap["GOOGLE_CREDENTIALS"] = string(data[:])
	} else if provider == AWS {
		providerDir = "aws/"

		envMap["AWS_ACCESS_KEY_ID"] = projectConfig.AWS_AccessKeyId
		envMap["AWS_SECRET_ACCESS_KEY"] = projectConfig.AWS_SecretAccessKey
	}

	stackDirectory := filepath.Join(ROOTDIR + providerDir + stackId)
	envvars := auto.EnvVars(envMap)

	// TODO: handle case for different providers
	cfg := auto.ConfigMap{
		"gcp:project": {Value: projectConfig.ProjectName},
		"aws:region":  {Value: "us-west-2"},
	}

	utils.Log(utils.DEBUG, "Checking resource path :"+stackDirectory)

	// TODO: if value updated, rewrite the yaml, or create a draft version
	_, err := os.Stat(stackDirectory + "/Pulumi.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			// File or directory does not exist
			_, err = os.Stat(stackDirectory)
			if err != nil {
				if os.IsNotExist(err) {
					if err := os.MkdirAll(stackDirectory, os.ModePerm); err != nil {
						panic(err)
					}
				}
			}

			os.WriteFile(stackDirectory+"/Pulumi.yaml", []byte(yamlConfigModel), 0644)
		} else {
			return nil, err
		}
	}

	stackName := auto.FullyQualifiedStackName("organization", projectConfig.ProjectName, stackId)
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, stackDirectory, envvars)
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
	// 	// 	return
	// }

	//utils.Log(utils.DEBUG, previewResult)

	stackUpResult, err := stack.Up(ctx)
	utils.Log(utils.DEBUG, "Creating Stack "+stackName)
	if err != nil {
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

// TODO: Currently we have hardcoded the logic for GCP, make this generic.
func PulumiStackDestroy(stackId string, projectConfig *ProjectData) {
	ctx := context.Background()
	stackDirectory := filepath.Join(ROOTDIR + "gcp/" + stackId)
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
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, stackDirectory, envvars)
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

// Sorts the resource config based on priority - VPC and Subnet should be deployed before other resources
//
// TODO: based on the resource dependency properly do the ordering, will require changes at he UI to add extra
// marker metadata to identify dependencies.
func getNextResourceIndexToDeploy(dataStore []DeploymentResource) *list.List {
	var first *list.Element
	var pendingIndexQueue *list.List = list.New()

	for i := 0; i < len(dataStore); i++ {
		if dataStore[i].ResourceType == COMPUTE_NETWORK || dataStore[i].ResourceType == EC2_VPC {
			if first == nil {
				first = pendingIndexQueue.PushFront(i)
			} else {
				pendingIndexQueue.PushFront(i)
			}
		} else if dataStore[i].ResourceType == COMPUTE_SUBNETWORK || dataStore[i].ResourceType == EC2_SUBNET {
			if first == nil {
				pendingIndexQueue.PushBack(i)
			} else {
				pendingIndexQueue.InsertBefore(i, first)
			}
		} else {
			utils.Log(utils.DEBUG, "Adding resource to queue: "+dataStore[i].Name)
			pendingIndexQueue.PushBack(i)
		}
	}

	return pendingIndexQueue
}

func SaveServeState() {
	data, _ := json.Marshal(backendGlobalContext)
	os.WriteFile("./.localStore/ServerState", data, 0644)
}
