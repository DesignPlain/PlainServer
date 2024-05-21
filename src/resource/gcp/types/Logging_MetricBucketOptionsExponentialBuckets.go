package types

type Logging_MetricBucketOptionsExponentialBuckets struct {
	// Must be greater than 0.
	Scale float64 `json:"scale,omitempty" yaml:"scale,omitempty"`

	// Must be greater than 1.
	GrowthFactor float64 `json:"growthFactor,omitempty" yaml:"growthFactor,omitempty"`

	// Must be greater than 0.
	NumFiniteBuckets int `json:"numFiniteBuckets,omitempty" yaml:"numFiniteBuckets,omitempty"`
}
