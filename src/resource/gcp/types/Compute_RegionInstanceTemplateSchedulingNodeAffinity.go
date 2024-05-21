package types

type Compute_RegionInstanceTemplateSchedulingNodeAffinity struct {
	/*
	   The operator. Can be `IN` for node-affinities
	   or `NOT_IN` for anti-affinities.
	*/
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	// Corresponds to the label values of a reservation resource.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// The key for the node affinity label.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
