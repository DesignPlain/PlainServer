package types

type Index_getRegionsFilter struct {
	// Name of the filter field. Valid values can be found in the [describe-regions AWS CLI Reference][1].
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
