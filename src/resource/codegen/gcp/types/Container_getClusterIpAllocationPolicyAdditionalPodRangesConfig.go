package types

type Container_getClusterIpAllocationPolicyAdditionalPodRangesConfig struct {
	// Name for pod secondary ipv4 range which has the actual range defined ahead.
	PodRangeNames []string `json:"podRangeNames,omitempty" yaml:"podRangeNames,omitempty"`
}
