package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type DeviceFleet struct {
	// Specifies details about the repository. see Output Config details below.
	OutputConfig types.Sagemaker_DeviceFleetOutputConfig `json:"outputConfig,omitempty" yaml:"outputConfig,omitempty"`

	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A description of the fleet.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the Device Fleet (must be unique).
	DeviceFleetName string `json:"deviceFleetName,omitempty" yaml:"deviceFleetName,omitempty"`

	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias bool `json:"enableIotRoleAlias,omitempty" yaml:"enableIotRoleAlias,omitempty"`
}
