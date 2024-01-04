package main

import (
	baseModel "DesignSphere_Server/src/resource"
	gcp "DesignSphere_Server/src/resource/GCP"
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
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
		c.String(http.StatusForbidden, "Please configure the cloud project")
		return
	}

	_controllerStatus.Count += 1
	var resourceOutputs []ResourceOutput

	// TODO: update the logic and datastructure to respect the resource dependency
	//       eg: VPCs should be created before the subnet and compute instances
	//       will have to add a DAG.
	if _controllerStatus.DataStore != nil {
		doDeploy := true
		var nextResourceIndex int
		pendingIndexQueue := getNextResourceIndexToDeploy(_controllerStatus.DataStore)
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
					var resourceConfigModel baseModel.ResourceModel

					if !slices.Contains(_controllerStatus.CurrentStack, data.Id.String()) {
						switch data.ResourceType {
						case Simple_Storage_Service:
							{
								resourceConfigModel = gcp.CreateBucketModel(
									projectConfig.ProjectName,
									data.Name,
									[]string{data.ResourceConfig["Members"].(string)},
									fmt.Sprintf(`"%s"`, data.ResourceConfig["Role"].(string)),
									data.ResourceConfig["Location"].(string))

								break
							}
						case EC2:
							{
								resourceConfigModel = gcp.CreateComputeInstanceModel(
									projectConfig.ProjectName,
									data.Name,
									data.ResourceConfig["Network"].(string),
									data.ResourceConfig["ServiceAccountEmail"].(string),
								)

								break
							}
						case Virtual_Private_Cloud:
							{
								// autoSubnet := data.ResourceConfig["AutoCreateSubNetwork"]
								mtu, _ := strconv.Atoi(data.ResourceConfig["MTU"].(string))
								resourceConfigModel = gcp.CreateVPCModel(
									projectConfig.ProjectName,
									data.Name,
									mtu,
									data.ResourceConfig["RoutingMode"].(string))

								break
							}
						case Subnet:
							{
								break
							}
						default:
							{
								c.String(http.StatusForbidden, "Please provide a valid GCP Cloud resource")
								break
							}
						}
					} else {
						Log(DEBUG, "Stack already contains the data")
					}

					if resourceConfigModel.Name != "" {
						serializedResourceConfigYaml, err := yaml.Marshal(resourceConfigModel)
						Log(DEBUG, string(serializedResourceConfigYaml))
						if err != nil {
							log.Fatal(err)
						}

						resourceOutputs, err = PulumiStackUp(data.Id.String(), string(serializedResourceConfigYaml), &projectConfig)

						if err == nil {
							_controllerStatus.CurrentStack = append(_controllerStatus.CurrentStack, data.Id.String())
							Log(DEBUG, "Adding new data to stack "+strings.Join(_controllerStatus.CurrentStack, " "))
						} else {
							Log(DEBUG, "Stack Up failed with error: ")
							Log(DEBUG, err)
						}
						data.Outputs = resourceOutputs
					}

					break
				}
			case AWS:
				{
				}
			case AZURE:
				{
				}
			default:
				c.String(http.StatusForbidden, "Please provide a valid Cloud provider")
			}
		}
	}

	Log(DEBUG, _controllerStatus.DataStore)
	Log(DEBUG, "Deploying stack")
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
	Log(DEBUG, "Deployment Plan is:")
	fmt.Println(pendingIndexQueue)
	return pendingIndexQueue
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
		c.String(http.StatusForbidden, "Please configure the cloud project")
		return
	}

	projectConfig := _controllerStatus.ProjectConfig["GCP"]
	Log(DEBUG, "Destroying stack")
	Log(DEBUG, _controllerStatus.CurrentStack)
	for i := 0; i < len(_controllerStatus.CurrentStack); i++ {
		Log(DEBUG, _controllerStatus.CurrentStack)
		PulumiStackDestroy(_controllerStatus.CurrentStack[i], &projectConfig)
	}

	// TODO: Check if any of the stack destroy logic failed
	_controllerStatus.CurrentStack = nil

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

func PulumiStackUp(stackId string, yamlConfigModel string, projectConfig *ProjectData) ([]ResourceOutput, error) {
	ctx := context.Background()
	// TODO: Make directory for different providers
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

	// TODO: handle case for different providers
	cfg := auto.ConfigMap{
		"gcp:project": {Value: projectConfig.ProjectName},
	}

	Log(DEBUG, workDir)

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

			f, err := os.Create(workDir + "/Pulumi.yaml")
			check(err)

			defer f.Close()

			_, err = f.WriteString(yamlConfigModel)
			check(err)

			f.Sync()
		} else {
			return nil, err
		}
	}

	stackName := auto.FullyQualifiedStackName("organization", projectConfig.ProjectName, stackId)
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
		Log(DEBUG, stackUpResult)
		return nil, err
	}
	Log(DEBUG, stackUpResult)
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

func PulumiStackDestroy(stackId string, projectConfig *ProjectData) {
	ctx := context.Background()
	workDir := filepath.Join(ROOTDIR + "gcp/" + stackId)
	data, err := os.ReadFile(ROOTDIR + projectConfig.APIKeyFileName)
	if err != nil {
		Log(ERROR, err)
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
