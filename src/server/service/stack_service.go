package service

import (
	. "DesignSphere_Server/src/server/model"
	"DesignSphere_Server/src/server/storage"
	"DesignSphere_Server/src/utils"
	"container/list"
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"gopkg.in/yaml.v2"
)

const ROOTDIR = "/Users/nmbr7/.NMBR7/Projects/DesignSphere/TestBase/"
const PULUMI_PROJECT_NAME = "DS_PulumiProject"

type ProjectConfig struct {
	GCP_ProjectName     string
	GCP_APIKeyFileName  string
	AWS_AccessKeyId     string
	AWS_SecretAccessKey string
}

type StackService struct {
	DataStore     storage.Store
	ProjectConfig ProjectConfig
}

func (stack_service *StackService) PulumiStackUp(provider ProviderType, stackId string, yamlConfigModel string) ([]ResourceOutput, error) {
	ctx := context.Background()
	projectConfig, _ := stack_service.GetProjectConfig()

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
		"gcp:project": {Value: projectConfig.GCP_ProjectName},
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

	stackName := auto.FullyQualifiedStackName("organization", PULUMI_PROJECT_NAME, stackId)
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
	if err != nil {
		rmErr := os.Remove(stackDirectory + "/Pulumi.yaml")
		if rmErr != nil {
			utils.Log(utils.ERROR, rmErr.Error())
		}
		return nil, err
	}

	utils.Log(utils.DEBUG, "Creating Stack "+stackName)
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
func (stack_service *StackService) PulumiStackDestroy(stackId string, provider ProviderType) error {
	ctx := context.Background()
	projectConfig, _ := stack_service.GetProjectConfig()

	var envMap = map[string]string{"PULUMI_CONFIG_PASSPHRASE": "test"}

	var providerDir string
	if provider == GCP {
		providerDir = "gcp/"

		data, err := os.ReadFile(ROOTDIR + projectConfig.GCP_APIKeyFileName)
		if err != nil {
			utils.Log(utils.ERROR, err)
			return err
		}

		envMap["GOOGLE_CREDENTIALS"] = string(data[:])
	} else if provider == AWS {
		providerDir = "aws/"

		envMap["AWS_ACCESS_KEY_ID"] = projectConfig.AWS_AccessKeyId
		envMap["AWS_SECRET_ACCESS_KEY"] = projectConfig.AWS_SecretAccessKey
	}

	stackDirectory := filepath.Join(ROOTDIR + providerDir + stackId)

	_, err := os.Stat(stackDirectory)
	if err == nil {
		envvars := auto.EnvVars(envMap)

		// TODO: handle case for different providers
		cfg := auto.ConfigMap{
			"gcp:project": {Value: projectConfig.GCP_ProjectName},
			"aws:region":  {Value: "us-west-2"},
		}

		stackName := auto.FullyQualifiedStackName("organization", PULUMI_PROJECT_NAME, stackId)
		stack, err := auto.UpsertStackLocalSource(ctx, stackName, stackDirectory, envvars)
		stack.SetAllConfig(ctx, cfg)

		if err != nil {
			utils.Log(utils.ERROR, err)
			return err
		}

		stackDestroyResult, err := stack.Destroy(ctx)
		utils.Log(utils.DEBUG, "Destroying Stack "+stackName)
		if err != nil {
			utils.Log(utils.ERROR, err)
			return err
		}
		_ = stack.Workspace().RemoveStack(ctx, stackName)
		utils.Log(utils.DEBUG, stackDestroyResult)

		return nil
	}

	utils.Log(utils.DEBUG, "Stack not found, might have been manually deleted")
	return nil
}

func (stack_service *StackService) GetProjectConfig() (*ProjectConfig, error) {
	var project_config ProjectConfig

	value, err := stack_service.DataStore.Get([]byte("ProjectConfig"))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(value, &project_config); err != nil {
		return nil, err
	}

	return &project_config, nil
}

func (stack_service *StackService) SetProjectConfig(project_config ProjectConfig) error {
	jsonValue, err := json.Marshal(project_config)
	if err != nil {
		return err
	}
	err = stack_service.DataStore.Set([]byte("ProjectConfig"), jsonValue)
	if err != nil {
		return err
	}

	return nil
}

func (stack_service *StackService) SaveState(resources []DeploymentResource) error {
	var resIdArray []string
	for _, res := range resources {
		// TODO: validate UUID
		resIdArray = append(resIdArray, res.Id.String())
		err := stack_service.saveResource(&res)
		if err != nil {
			return err
		}
	}

	jsonArrayValue, err := json.Marshal(resIdArray)
	if err != nil {
		return err
	}

	if err := stack_service.DataStore.Set([]byte("ResourceIds"), jsonArrayValue); err != nil {
		return err
	}

	return nil
}

func (stack_service *StackService) GetState() ([]DeploymentResource, error) {
	var resourceIdArray []string

	value, err := stack_service.DataStore.Get([]byte("ResourceIds"))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(value, &resourceIdArray); err != nil {
		return nil, err
	}

	var resources []DeploymentResource
	for _, resId := range resourceIdArray {
		res, err := stack_service.getResourceById(resId)
		if err != nil {
			return nil, err
		}

		resources = append(resources, *res)
	}

	return resources, nil
}

func (stack_service *StackService) getResourceById(resId string) (*DeploymentResource, error) {
	var resource DeploymentResource
	value, err := stack_service.DataStore.Get([]byte(resId))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(value, &resource); err != nil {
		return nil, err
	}

	return &resource, nil
}

func (stack_service *StackService) saveResource(resource *DeploymentResource) error {
	jsonValue, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	if err := stack_service.DataStore.Set([]byte(resource.Id.String()), jsonValue); err != nil {
		return err
	}

	return nil
}

func (stack_service *StackService) DeployResources() error {

	// if !ValidateState() {
	// 	c.String(http.StatusForbidden, "Please configure the cloud project")
	// 	return
	// }

	_, err := stack_service.GetProjectConfig()
	if err != nil {
		// TODO: return proper error
		return err
	}

	resources, err := stack_service.GetState()
	if err != nil {
		// TODO: check if the resources are available
		return err
	}

	var resourceOutputs []ResourceOutput

	// TODO: update the logic and datastructure to respect the resource dependency
	//       eg: VPCs should be created before the subnet and compute instances
	//       will have to add a DAG.
	if resources != nil {
		doDeploy := true
		var nextResourceIndex int
		//var deploymentList []int

		pendingIndexQueue := getNextResourceIndexToDeploy(resources)
		// for i := pendingIndexQueue.Front(); i != nil; i = i.Next() {
		// 	deploymentList = append(deploymentList, i.Value.(int))
		// }

		stackDeploymentOrder, err := stack_service.getStackDeployOrder()
		if err != nil {
			stackDeploymentOrder = make([]string, 0)
		}

		for doDeploy {
			if pendingIndexQueue.Len() > 0 {
				front := pendingIndexQueue.Front()
				nextResourceIndex = front.Value.(int)
				pendingIndexQueue.Remove(front)
			} else {
				doDeploy = false
				break
			}
			data := &resources[nextResourceIndex]
			utils.Log(utils.DEBUG, "Processing resource: "+data.Id.String()+" "+data.Name)
			switch data.ProviderType {
			case GCP:
				{
					var resourceConfigModel ResourceModel
					if !slices.Contains(stackDeploymentOrder, data.Id.String()) {
						resourceConfigModel = CreateResourceModel(PULUMI_PROJECT_NAME, data.ResourceType.String(), data.ResourceType, data.Name, data.ResourceConfig)

					} else {
						utils.Log(utils.DEBUG, "Stack already contains the data")
					}

					if resourceConfigModel.Name != "" {
						serializedResourceConfigYaml, err := yaml.Marshal(resourceConfigModel)
						utils.Log(utils.DEBUG, string(serializedResourceConfigYaml))
						if err != nil {
							log.Fatal(err)
							continue
						}

						resourceOutputs, err = stack_service.PulumiStackUp(GCP, data.Id.String(), string(serializedResourceConfigYaml))

						if err == nil {
							data.Status = Deployed
							stackDeploymentOrder = append(stackDeploymentOrder, data.Id.String())
							utils.Log(utils.DEBUG, "Adding new data to stack "+strings.Join(stackDeploymentOrder, " "))

						} else {
							data.Status = Errored
							utils.Log(utils.ERROR, "Stack Up failed with error: "+err.Error())
						}

						resources[nextResourceIndex] = *data
						data.Outputs = resourceOutputs
					}
				}
			case AWS:
				{
					var resourceConfigModel ResourceModel
					if !slices.Contains(stackDeploymentOrder, data.Id.String()) {
						resourceConfigModel = CreateResourceModel(PULUMI_PROJECT_NAME, data.ResourceType.String(), data.ResourceType, data.Name, data.ResourceConfig)
					} else {
						utils.Log(utils.DEBUG, "Stack already contains the data")
					}

					if resourceConfigModel.Resources != nil {
						serializedResourceConfigYaml, err := yaml.Marshal(resourceConfigModel)
						utils.Log(utils.DEBUG, string(serializedResourceConfigYaml))
						if err != nil {
							log.Fatal(err)
							continue
						}

						resourceOutputs, err = stack_service.PulumiStackUp(AWS, data.Id.String(), string(serializedResourceConfigYaml))

						if err == nil {
							data.Status = Deployed
							stackDeploymentOrder = append(stackDeploymentOrder, data.Id.String())
							utils.Log(utils.DEBUG, "Adding new data to stack "+strings.Join(stackDeploymentOrder, " "))
						} else {
							data.Status = Errored
							utils.Log(utils.ERROR, "Stack Up failed with error: "+err.Error())
						}

						resources[nextResourceIndex] = *data
						data.Outputs = resourceOutputs
					}
				}
			// case AZURE:
			// 	{
			// 	}
			default:
				return UnknownResourceProvider
			}
		}

		if err := stack_service.saveStackDeploymentOrder(stackDeploymentOrder); err != nil {
			return err
		}
	}

	if err := stack_service.SaveState(resources); err != nil {
		return err
	}

	resources, _ = stack_service.GetState()
	var resString, _ = json.Marshal(resources)
	utils.Log(utils.DEBUG, "DeployedResources from store "+string(resString))

	return nil
}

func (stack_service *StackService) DestroyResources() error {

	// if !_controllerStatus.ValidateState() {
	// 	c.String(http.StatusForbidden, "Please configure the cloud project")
	// 	return
	// }

	_, err := stack_service.GetProjectConfig()
	if err != nil {
		// TODO: return proper error
		return err
	}

	stackDeploymentOrder, err := stack_service.getStackDeployOrder()
	if err != nil {
		return NoResourcesAvailable
	}

	utils.Log(utils.DEBUG, "Destroying stack with order: ")
	utils.Log(utils.DEBUG, stackDeploymentOrder)
	tempStackOrder := make([]string, 0)
	for _, resId := range stackDeploymentOrder {
		resource, err := stack_service.getResourceById(resId)
		if err != nil {
			utils.Log(utils.ERROR, "Getting resource return error:  "+err.Error())
			continue
		}

		err = stack_service.PulumiStackDestroy(resId, resource.ProviderType)

		if err != nil {
			utils.Log(utils.ERROR, "Stack destroy failed with error: "+err.Error())
			tempStackOrder = append(tempStackOrder, resId)
			continue
		}

		resource.Status = Draft

		if err = stack_service.saveResource(resource); err != nil {
			utils.Log(utils.ERROR, "Saving resource return error:  "+err.Error())
		}
	}

	if err = stack_service.saveStackDeploymentOrder(tempStackOrder); err != nil {
		return err
	}

	return nil
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

func (stack_service *StackService) getStackDeployOrder() ([]string, error) {
	var current_stack_order []string
	value, err := stack_service.DataStore.Get([]byte("StackDeploymentOrder"))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(value, &current_stack_order); err != nil {
		return nil, err
	}

	return current_stack_order, nil
}

func (stack_service *StackService) saveStackDeploymentOrder(current_stack_order []string) error {
	jsonArrayValue, err := json.Marshal(current_stack_order)
	if err != nil {
		return err
	}

	if err := stack_service.DataStore.Set([]byte("StackDeploymentOrder"), jsonArrayValue); err != nil {
		return err
	}

	return nil
}
