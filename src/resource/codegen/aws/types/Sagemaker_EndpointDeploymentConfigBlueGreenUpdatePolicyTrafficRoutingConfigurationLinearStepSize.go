package types

type Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSize struct {
	// Specifies the endpoint capacity type. Valid values are: `INSTANCE_COUNT`, or `CAPACITY_PERCENT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
