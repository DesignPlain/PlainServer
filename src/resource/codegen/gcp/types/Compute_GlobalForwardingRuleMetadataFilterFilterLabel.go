package types

type Compute_GlobalForwardingRuleMetadataFilterFilterLabel struct {
	/*
	   Name of the metadata label. The length must be between
	   1 and 1024 characters, inclusive.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The value that the label must match. The value has a maximum
	   length of 1024 characters.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
