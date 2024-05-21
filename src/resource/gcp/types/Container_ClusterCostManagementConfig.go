package types

type Container_ClusterCostManagementConfig struct {
	// Whether to enable the [cost allocation](https://cloud.google.com/kubernetes-engine/docs/how-to/cost-allocations) feature.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
