package types

type Compute_RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel struct {
	/*
	   Name of metadata label. The name can have a maximum length of 1024 characters
	   and must be at least 1 character long.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The value of the label must match the specified value. value can have a maximum
	   length of 1024 characters.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
