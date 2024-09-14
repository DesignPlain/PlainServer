package types

type Container_ClusterNodePoolNodeConfigSoleTenantConfigNodeAffinity struct {
	// The default or custom node affinity label key name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Specifies affinity or anti-affinity. Accepted values are `"IN"` or `"NOT_IN"`
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	// List of node affinity label values as strings.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
