package types

type Gkehub_FleetDefaultClusterConfig struct {
	/*
	   Enable/Disable binary authorization features for the cluster.
	   Structure is documented below.
	*/
	BinaryAuthorizationConfig Gkehub_FleetDefaultClusterConfigBinaryAuthorizationConfig `json:"binaryAuthorizationConfig,omitempty" yaml:"binaryAuthorizationConfig,omitempty"`

	/*
	   Enable/Disable Security Posture features for the cluster.
	   Structure is documented below.
	*/
	SecurityPostureConfig Gkehub_FleetDefaultClusterConfigSecurityPostureConfig `json:"securityPostureConfig,omitempty" yaml:"securityPostureConfig,omitempty"`
}
