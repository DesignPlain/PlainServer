package types

type Wafv2_WebAclLoggingConfigurationLoggingFilter struct {
	// Default handling for logs that don't match any of the specified filtering conditions. Valid values for `default_behavior` are `KEEP` or `DROP`.
	DefaultBehavior string `json:"defaultBehavior,omitempty" yaml:"defaultBehavior,omitempty"`

	// Filter(s) that you want to apply to the logs. See Filter below for more details.
	Filters []Wafv2_WebAclLoggingConfigurationLoggingFilterFilter `json:"filters,omitempty" yaml:"filters,omitempty"`
}
