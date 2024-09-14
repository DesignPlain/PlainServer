package types

type Logging_MetricBucketOptionsLinearBuckets struct {
	// Must be greater than 0.
	NumFiniteBuckets int `json:"numFiniteBuckets,omitempty" yaml:"numFiniteBuckets,omitempty"`

	// Lower bound of the first bucket.
	Offset float64 `json:"offset,omitempty" yaml:"offset,omitempty"`

	// Must be greater than 0.
	Width float64 `json:"width,omitempty" yaml:"width,omitempty"`
}
