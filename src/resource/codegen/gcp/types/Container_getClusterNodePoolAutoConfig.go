package types

type Container_getClusterNodePoolAutoConfig struct {
	// Collection of Compute Engine network tags that can be applied to a node's underlying VM instance.
	NetworkTags []Container_getClusterNodePoolAutoConfigNetworkTag `json:"networkTags,omitempty" yaml:"networkTags,omitempty"`
}
