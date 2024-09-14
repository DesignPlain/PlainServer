package types

type Logging_MetricBucketOptions struct {
	/*
	   Specifies a set of buckets with arbitrary widths.
	   Structure is documented below.
	*/
	ExplicitBuckets Logging_MetricBucketOptionsExplicitBuckets `json:"explicitBuckets,omitempty" yaml:"explicitBuckets,omitempty"`

	/*
	   Specifies an exponential sequence of buckets that have a width that is proportional to the value of
	   the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.
	   Structure is documented below.
	*/
	ExponentialBuckets Logging_MetricBucketOptionsExponentialBuckets `json:"exponentialBuckets,omitempty" yaml:"exponentialBuckets,omitempty"`

	/*
	   Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
	   Each bucket represents a constant absolute uncertainty on the specific value in the bucket.
	   Structure is documented below.
	*/
	LinearBuckets Logging_MetricBucketOptionsLinearBuckets `json:"linearBuckets,omitempty" yaml:"linearBuckets,omitempty"`
}
