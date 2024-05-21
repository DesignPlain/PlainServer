package logging

type OrganizationExclusion struct {
	// A human-readable description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Whether this exclusion rule should be disabled or not. This defaults to
	   false.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	   See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	   write a filter.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// The name of the logging exclusion.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The organization to create the exclusion in.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`
}
