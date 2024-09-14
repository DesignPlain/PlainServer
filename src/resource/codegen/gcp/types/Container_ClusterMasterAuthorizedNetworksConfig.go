package types

type Container_ClusterMasterAuthorizedNetworksConfig struct {
	/*
	   Whether Kubernetes master is
	   accessible via Google Compute Engine Public IPs.
	*/
	GcpPublicCidrsAccessEnabled bool `json:"gcpPublicCidrsAccessEnabled,omitempty" yaml:"gcpPublicCidrsAccessEnabled,omitempty"`

	/*
	   External networks that can access the
	   Kubernetes cluster master through HTTPS.
	*/
	CidrBlocks []Container_ClusterMasterAuthorizedNetworksConfigCidrBlock `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`
}
