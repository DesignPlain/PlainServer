package types

type Compute_InstanceSchedulingNodeAffinity struct {
	// The key for the node affinity label.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   The operator. Can be `IN` for node-affinities
	   or `NOT_IN` for anti-affinities.
	*/
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	// The values for the node affinity label.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
