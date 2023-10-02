package main

type Any interface{}

type DeploymentResource struct {
	Name           string       `json:"title"`
	ResourceType   ResourceType `json:"resourceType"`
	ProviderType   ProviderType `json:"platformType"`
	Shape          UIShape      `json:"shape"`
	Position       UIPosition   `json:"position"`
	ResourceConfig any          `json:"resourceConfig"`
}

type UIPosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type UIShape struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

type GCPBucket struct {
	Location                 string `json:"location"`
	UniformBucketLevelAccess bool   `json:"uniformBucketLevelAccess"`
	Role                     string `json:"role"`
	Members                  string `json:"members"`
}

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

type ProviderType int64

const (
	GCP ProviderType = iota
	AWS
	AZURE
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
