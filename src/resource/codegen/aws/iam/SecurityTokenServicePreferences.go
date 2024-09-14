package iam

type SecurityTokenServicePreferences struct {
	// The version of the STS global endpoint token. Valid values: `v1Token`, `v2Token`.
	GlobalEndpointTokenVersion string `json:"globalEndpointTokenVersion,omitempty" yaml:"globalEndpointTokenVersion,omitempty"`
}
