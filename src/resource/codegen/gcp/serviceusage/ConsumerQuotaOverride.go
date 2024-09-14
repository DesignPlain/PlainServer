package serviceusage

type ConsumerQuotaOverride struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions map[string]string `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	/*
	   If the new quota would decrease the existing quota by more than 10%!!(MISSING),(MISSING) the request is rejected.
	   If `force` is `true`, that safety check is ignored.
	*/
	Force bool `json:"force,omitempty" yaml:"force,omitempty"`

	/*
	   The limit on the metric, e.g. `/project/region`.
	   > Make sure that `limit` is in a format that doesn't start with `1/` or contain curly braces.
	   E.g. use `/project/user` instead of `1/{project}/{user}`.


	   - - -
	*/
	Limit string `json:"limit,omitempty" yaml:"limit,omitempty"`

	// The metric that should be limited, e.g. `compute.googleapis.com/cpus`.
	Metric string `json:"metric,omitempty" yaml:"metric,omitempty"`

	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue string `json:"overrideValue,omitempty" yaml:"overrideValue,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The service that the metrics belong to, e.g. `compute.googleapis.com`.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
