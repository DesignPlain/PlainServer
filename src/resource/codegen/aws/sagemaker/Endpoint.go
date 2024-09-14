package sagemaker

import types "libds/aws/types"

type Endpoint struct {
	// The name of the endpoint configuration to use.
	EndpointConfigName string `json:"endpointConfigName,omitempty" yaml:"endpointConfigName,omitempty"`

	// The name of the endpoint. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations. See Deployment Config.
	DeploymentConfig types.Sagemaker_EndpointDeploymentConfig `json:"deploymentConfig,omitempty" yaml:"deploymentConfig,omitempty"`
}
