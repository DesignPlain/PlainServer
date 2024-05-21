package types

type Gkehub_FleetDefaultClusterConfigSecurityPostureConfig struct {
	/*
	   Sets which mode to use for Security Posture features.
	   Possible values are: `DISABLED`, `BASIC`.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   Sets which mode to use for vulnerability scanning.
	   Possible values are: `VULNERABILITY_DISABLED`, `VULNERABILITY_BASIC`, `VULNERABILITY_ENTERPRISE`.
	*/
	VulnerabilityMode string `json:"vulnerabilityMode,omitempty" yaml:"vulnerabilityMode,omitempty"`
}
