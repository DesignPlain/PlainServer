package types

type Compute_getGlobalForwardingRuleMetadataFilterFilterLabel struct {
	/*
	   The name of the global forwarding rule.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The value that the label must match. The value has a maximum
	   length of 1024 characters.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
