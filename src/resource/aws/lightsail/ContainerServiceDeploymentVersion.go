package lightsail

import types "DesignSphere_Server/src/resource/aws/types"

type ContainerServiceDeploymentVersion struct {
	// A set of configuration blocks that describe the settings of the containers that will be launched on the container service. Maximum of 53. Detailed below.
	Containers []types.Lightsail_ContainerServiceDeploymentVersionContainer `json:"containers,omitempty" yaml:"containers,omitempty"`

	// A configuration block that describes the settings of the public endpoint for the container service. Detailed below.
	PublicEndpoint types.Lightsail_ContainerServiceDeploymentVersionPublicEndpoint `json:"publicEndpoint,omitempty" yaml:"publicEndpoint,omitempty"`

	// The name for the container service.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
}
