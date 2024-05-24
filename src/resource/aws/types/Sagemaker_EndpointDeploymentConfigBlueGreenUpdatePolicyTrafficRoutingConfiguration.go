package types

type Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration struct {
	// The waiting time (in seconds) between incremental steps to turn on traffic on the new endpoint fleet. Valid values are between `0` and `3600`.
	WaitIntervalInSeconds int `json:"waitIntervalInSeconds,omitempty" yaml:"waitIntervalInSeconds,omitempty"`

	// Batch size for the first step to turn on traffic on the new endpoint fleet. Value must be less than or equal to 50%!o(MISSING)f the variant's total instance count. See Canary Size.
	CanarySize Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySize `json:"canarySize,omitempty" yaml:"canarySize,omitempty"`

	// Batch size for each step to turn on traffic on the new endpoint fleet. Value must be 10-50%!o(MISSING)f the variant's total instance count. See Linear Step Size.
	LinearStepSize Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSize `json:"linearStepSize,omitempty" yaml:"linearStepSize,omitempty"`

	// Traffic routing strategy type. Valid values are: `ALL_AT_ONCE`, `CANARY`, and `LINEAR`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
