package types

type Compute_getInstanceTemplateSchedulingNodeAffinity struct {
	/*
	   The operator. Can be `IN` for node-affinities
	   or `NOT_IN` for anti-affinities.
	*/
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// The key for the node affinity label.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
