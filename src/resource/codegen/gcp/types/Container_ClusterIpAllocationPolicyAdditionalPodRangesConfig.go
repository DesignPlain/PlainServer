package types

type Container_ClusterIpAllocationPolicyAdditionalPodRangesConfig struct {
	// The names of the Pod ranges to add to the cluster.
	PodRangeNames []string `json:"podRangeNames,omitempty" yaml:"podRangeNames,omitempty"`
}
