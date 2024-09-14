package types

type Container_ClusterNodePoolAutoConfig struct {
	// The network tag config for the cluster's automatically provisioned node pools.
	NetworkTags Container_ClusterNodePoolAutoConfigNetworkTags `json:"networkTags,omitempty" yaml:"networkTags,omitempty"`
}
