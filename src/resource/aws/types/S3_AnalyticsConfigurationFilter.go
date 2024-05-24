package types

type S3_AnalyticsConfigurationFilter struct {
	// Object prefix for filtering.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Set of object tags for filtering.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
