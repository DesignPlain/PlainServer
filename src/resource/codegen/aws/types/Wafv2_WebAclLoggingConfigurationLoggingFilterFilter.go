package types

type Wafv2_WebAclLoggingConfigurationLoggingFilterFilter struct {
	// Parameter that determines how to handle logs that meet the conditions and requirements of the filter. The valid values for `behavior` are `KEEP` or `DROP`.
	Behavior string `json:"behavior,omitempty" yaml:"behavior,omitempty"`

	// Match condition(s) for the filter. See Condition below for more details.
	Conditions []Wafv2_WebAclLoggingConfigurationLoggingFilterFilterCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// Logic to apply to the filtering conditions. You can specify that a log must match all conditions or at least one condition in order to satisfy the filter. Valid values for `requirement` are `MEETS_ALL` or `MEETS_ANY`.
	Requirement string `json:"requirement,omitempty" yaml:"requirement,omitempty"`
}
