package types

type Container_getClusterNodePoolQueuedProvisioning struct {
	// Whether nodes in this node pool are obtainable solely through the ProvisioningRequest API
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
