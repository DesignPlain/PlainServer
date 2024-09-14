package types

type Container_ClusterSecurityPostureConfig struct {
	// Sets the mode of the Kubernetes security posture API's off-cluster features. Available options include `DISABLED` and `BASIC`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Available options include `VULNERABILITY_DISABLED`, `VULNERABILITY_BASIC` and `VULNERABILITY_ENTERPRISE`.
	VulnerabilityMode string `json:"vulnerabilityMode,omitempty" yaml:"vulnerabilityMode,omitempty"`
}
