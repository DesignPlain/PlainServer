package types

type Cloudfront_ContinuousDeploymentPolicyTrafficConfigSingleWeightConfig struct {
	// Session stickiness provides the ability to define multiple requests from a single viewer as a single session. This prevents the potentially inconsistent experience of sending some of a given user's requests to the staging distribution, while others are sent to the primary distribution. Define the session duration using TTL values. See `session_stickiness_config`.
	SessionStickinessConfig Cloudfront_ContinuousDeploymentPolicyTrafficConfigSingleWeightConfigSessionStickinessConfig `json:"sessionStickinessConfig,omitempty" yaml:"sessionStickinessConfig,omitempty"`

	// The percentage of traffic to send to a staging distribution, expressed as a decimal number between `0` and `.15`.
	Weight float64 `json:"weight,omitempty" yaml:"weight,omitempty"`
}
