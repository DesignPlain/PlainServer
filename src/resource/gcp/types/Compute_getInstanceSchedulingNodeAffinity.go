package types

type Compute_getInstanceSchedulingNodeAffinity struct {
	//
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	//
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
