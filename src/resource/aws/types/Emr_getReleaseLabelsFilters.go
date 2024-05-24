package types

type Emr_getReleaseLabelsFilters struct {
	// Optional release label application filter. For example, `Spark@2.1.0` or `Spark`.
	Application string `json:"application,omitempty" yaml:"application,omitempty"`

	// Optional release label version prefix filter. For example, `emr-5`.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
