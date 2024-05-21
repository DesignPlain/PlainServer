package types

type Logging_getSinkExclusion struct {
	// An advanced logs filter that matches the log entries to be excluded.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// A client-assigned identifier, such as `load-balancer-exclusion`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A description of this exclusion.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether this exclusion is disabled and it does not exclude any log entries.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}
