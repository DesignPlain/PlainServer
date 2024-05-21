package types

type Logging_MetricBucketOptionsExplicitBuckets struct {
	// The values must be monotonically increasing.
	Bounds []float64 `json:"bounds,omitempty" yaml:"bounds,omitempty"`
}
