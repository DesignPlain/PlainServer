package model

import (
	"DesignSphere_Server/src/resource/aws/ec2"
	"DesignSphere_Server/src/resource/aws/elb"
	"DesignSphere_Server/src/resource/aws/lambda"
	"DesignSphere_Server/src/resource/aws/s3"

	"DesignSphere_Server/src/resource/gcp/cloudfunctionsv2"
	"DesignSphere_Server/src/resource/gcp/compute"
	"DesignSphere_Server/src/resource/gcp/storage"

	"DesignSphere_Server/src/utils"
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
	IconSrc   string         `json:"iconSrc"`
	Status    ResourceStatus `json:"status"`
	Position  UIPosition     `json:"position"`
	Inlets    []string       `json:"inlets"`
	Outlets   []string       `json:"outlets"`
	InletMap  Any            `json:"inletMap"`
	OutletMap Any            `json:"outletMap"`

	/*
	   The configuration representing a cloud resource
	*/
	ResourceConfig Any              `json:"resourceConfig"`
	Outputs        []ResourceOutput `json:"resOutputs"`
}

/*
Custom JSON deserializer for the `DeploymentResource` type.
This method will deserialize `ResourceConfig` to it's respective concrete type based on the `ProviderType` and `Resourcetype`
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
	utils.Log(utils.DEBUG, string(rawResConfig))
	resConfig, err := GetResourceConfigInstance(c.ProviderType, c.ResourceType, rawResConfig)
	if err != nil {
		return err
	}

	c.ResourceConfig = resConfig

	return nil
}

var resourceTypeMap = map[ResourceType]func() Any{
	// GCP Resources
	COMPUTE_INSTANCE: func() Any {
		return &compute.Instance{}
	},
	COMPUTE_NETWORK: func() Any {
		return &compute.Network{}
	},
	STORAGE_BUCKET: func() Any {
		return &storage.Bucket{}
	},
	COMPUTE_SUBNETWORK: func() Any {
		return &compute.Subnetwork{}
	},
	CLOUDFUNCTIONSV2_FUNCTION: func() Any {
		return &cloudfunctionsv2.Function{}
	},
	COMPUTE_FORWARDINGRULE: func() Any {
		return &compute.ForwardingRule{}
	},

	// AWS Resources
	EC2_INSTANCE: func() Any {
		return &ec2.Instance{}
	},
	EC2_VPC: func() Any {
		return &ec2.Vpc{}
	},
	EC2_SUBNET: func() Any {
		return &ec2.Subnet{}
	},
	S3_BUCKET: func() Any {
		return &s3.Bucket{}
	},
	LAMBDA_FUNCTION: func() Any {
		return &lambda.Function{}
	},
	ELB_LOADBALANCER: func() Any {
		return &elb.LoadBalancer{}
	},
}

func GetResourceConfigInstance(provider_type ProviderType, resource_type ResourceType, rawResConfig json.RawMessage) (Any, error) {
	initResourceConfigObject, ok := resourceTypeMap[resource_type]

	if !ok {
		return nil, ResourceUnMarshalFailed
	}

	resConfig := initResourceConfigObject()
	if err := json.Unmarshal(rawResConfig, &resConfig); err != nil {
		return nil, ResourceUnMarshalFailed
	}

	// TODO: Validate the resource instance

	return resConfig, nil
}

/*
Server Error Type enum
*/
type ServerError int64

const (
	ResourceUnMarshalFailed ServerError = iota
	UnknownResourceProvider
	NoResourcesAvailable
)

func (err ServerError) Error() string {
	switch err {
	case ResourceUnMarshalFailed:
		return "Unmarshalling resource type failed."
	case UnknownResourceProvider:
		return "Unknown resource provider specified."
	case NoResourcesAvailable:
		return "No saved resources are available."
	}

	return "Unknown Server Error"
}

type ResourceOutput struct {
	Name  string `json:"name"`
	Value Any    `json:"value"`
}

type ResourceStatus int8

const (
	Draft ResourceStatus = iota
	Deployed
	Errored
)

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
Provider Type Enum for a cloud resource
*/
type ProviderType int64

const (
	UNKNOWN ProviderType = iota
	GCP
	AWS
	// AZURE
)

func (s ProviderType) String() string {
	switch s {
	case GCP:
		return "gcp"
	case AWS:
		return "aws"
	}

	return "unknown"
}
