package types

type Cloudfront_ContinuousDeploymentPolicyTrafficConfig struct {
	// Determines which HTTP requests are sent to the staging distribution. See `single_header_config`.
	SingleHeaderConfig Cloudfront_ContinuousDeploymentPolicyTrafficConfigSingleHeaderConfig `json:"singleHeaderConfig,omitempty" yaml:"singleHeaderConfig,omitempty"`

	// Contains the percentage of traffic to send to the staging distribution. See `single_weight_config`.
	SingleWeightConfig Cloudfront_ContinuousDeploymentPolicyTrafficConfigSingleWeightConfig `json:"singleWeightConfig,omitempty" yaml:"singleWeightConfig,omitempty"`

	// Type of traffic configuration. Valid values are `SingleWeight` and `SingleHeader`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
