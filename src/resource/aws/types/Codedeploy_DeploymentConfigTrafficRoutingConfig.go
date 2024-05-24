package types

type Codedeploy_DeploymentConfigTrafficRoutingConfig struct {
	// The time based canary configuration information. If `type` is `TimeBasedLinear`, use `time_based_linear` instead.
	TimeBasedCanary Codedeploy_DeploymentConfigTrafficRoutingConfigTimeBasedCanary `json:"timeBasedCanary,omitempty" yaml:"timeBasedCanary,omitempty"`

	// The time based linear configuration information. If `type` is `TimeBasedCanary`, use `time_based_canary` instead.
	TimeBasedLinear Codedeploy_DeploymentConfigTrafficRoutingConfigTimeBasedLinear `json:"timeBasedLinear,omitempty" yaml:"timeBasedLinear,omitempty"`

	// Type of traffic routing config. One of `TimeBasedCanary`, `TimeBasedLinear`, `AllAtOnce`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
