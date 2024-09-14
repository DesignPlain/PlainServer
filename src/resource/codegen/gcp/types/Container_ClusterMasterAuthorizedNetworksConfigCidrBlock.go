package types

type Container_ClusterMasterAuthorizedNetworksConfigCidrBlock struct {
	/*
	   External network that can access Kubernetes master through HTTPS.
	   Must be specified in CIDR notation.
	*/
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// Field for users to identify CIDR blocks.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
