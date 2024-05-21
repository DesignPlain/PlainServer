package types

type Container_getClusterMasterAuthorizedNetworksConfig struct {
	// External networks that can access the Kubernetes cluster master through HTTPS.
	CidrBlocks []Container_getClusterMasterAuthorizedNetworksConfigCidrBlock `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`

	// Whether master is accessbile via Google Compute Engine Public IP addresses.
	GcpPublicCidrsAccessEnabled bool `json:"gcpPublicCidrsAccessEnabled,omitempty" yaml:"gcpPublicCidrsAccessEnabled,omitempty"`
}
