package types

type Cloudfront_ContinuousDeploymentPolicyTrafficConfigSingleHeaderConfig struct {
	// Request header name to send to the staging distribution. The header must contain the prefix `aws-cf-cd-`.
	Header string `json:"header,omitempty" yaml:"header,omitempty"`

	// Request header value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
