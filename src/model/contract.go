package model

import (
	gcp "DesignSphere_Server/src/resource/gcp/compute"
	utils "DesignSphere_Server/src/utils"
	"encoding/json"

	"github.com/google/uuid"
)

type Any interface{}

/*
Deployment Resource model for cloud resource
*/
type DeploymentResource struct {
	Id       uuid.UUID `json:"id"`
	TypeName string    `json:"name"`
	Name     string    `json:"title"`

	/*
	   Resource type for respective cloud provider resource
	*/
	ResourceType ResourceType `json:"resourceType"`

	/*
	   Provider Type for the resource `[AWS,GCP,AZURE]`
	*/
	ProviderType ProviderType `json:"providerType"`
	Shape        UIShape      `json:"shape"`

	/*
	   Icon source location for the UI component icon
	*/
	IconSrc         string     `json:"iconSrc"`
	Position        UIPosition `json:"position"`
	Inlets          []string   `json:"inlets"`
	Outlets         []string   `json:"outlets"`
	InletMapString  string     `json:"inletMapString"`
	OutletMapString string     `json:"outletMapString"`

	/*
	   The configuration representing a cloud resource
	*/
	ResourceConfig Any              `json:"resourceConfig"`
	Outputs        []ResourceOutput `json:"resOutputs"`
}

/*
Custom JSON deserializer for the `DeploymentResource` type.
This method will deserialize `ResourceConfig` to it's respective concrete type based on the  `ProviderType` and `Resourcetype`
*/
func (c *DeploymentResource) UnmarshalJSON(data []byte) error {
	type DeploymentResourceAlias DeploymentResource
	var rawResConfig json.RawMessage

	tmpResource := struct {
		*DeploymentResourceAlias
	}{
		DeploymentResourceAlias: (*DeploymentResourceAlias)(c),
	}

	tmpResource.ResourceConfig = &rawResConfig

	if err := json.Unmarshal(data, &tmpResource); err != nil {
		return err
	}

	utils.Log(utils.DEBUG, tmpResource)

	utils.Log(utils.DEBUG, string(rawResConfig))

	resConfig, err := GetResourceConfigInstance(c.ProviderType, c.ResourceType, rawResConfig)
	if err != nil {
		return err
	}

	c.ResourceConfig = resConfig

	return nil
}

func GetResourceConfigInstance(provider_type ProviderType, resource_type ResourceType, rawResConfig json.RawMessage) (Any, error) {

	var resConfig gcp.Instance
	if err := json.Unmarshal(rawResConfig, &resConfig); err != nil {
		return nil, err
	}
	return resConfig, nil
}

type ResourceOutput struct {
	Name  string `json:"name"`
	Value Any    `json:"value"`
}

/*
UI position model for UI component rendering
*/
type UIPosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

/*
UI shape model for UI component rendering
*/
type UIShape struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

/*
Resource Type Enum for a cloud provider resource
*/
type ResourceType int64

const (
	Lambda ResourceType = iota
	Database
	EC2
	Simple_Storage_Service
	SQS
	Elastic_Load_Balancer
	APIGateway
	Virtual_Private_Cloud
	Subnet
	DynamoDB
)

func (s ResourceType) String() string {
	switch s {
	case Lambda:
		return "Lambda"
	case EC2:
		return "EC2"
	case Database:
		return "Database"
	case Virtual_Private_Cloud:
		return "Virtual_Private_Cloud"
	case Simple_Storage_Service:
		return "Simple_Storage_Service"
	}
	return "unknown"
}

/*
Provider Type Enum for a cloud resource
*/
type ProviderType int64

const (
	GCP ProviderType = iota
	AZURE
	AWS
)
