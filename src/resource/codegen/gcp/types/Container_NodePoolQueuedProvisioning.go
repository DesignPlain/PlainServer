package types

type Container_NodePoolQueuedProvisioning struct {
	// Makes nodes obtainable through the [ProvisioningRequest API](https://cloud.google.com/kubernetes-engine/docs/how-to/provisioningrequest) exclusively.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
