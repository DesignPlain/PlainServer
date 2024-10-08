package types

type Logging_OrganizationSinkExclusion struct {
	// A description of this exclusion.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100%!!(MISSING)o(MISSING)f the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	   write a filter.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
