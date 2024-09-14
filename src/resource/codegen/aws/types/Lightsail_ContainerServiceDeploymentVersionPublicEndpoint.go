package types

type Lightsail_ContainerServiceDeploymentVersionPublicEndpoint struct {
	// The port of the container to which traffic is forwarded to.
	ContainerPort int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`

	// A configuration block that describes the health check configuration of the container. Detailed below.
	HealthCheck Lightsail_ContainerServiceDeploymentVersionPublicEndpointHealthCheck `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	// The name of the container for the endpoint.
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`
}
