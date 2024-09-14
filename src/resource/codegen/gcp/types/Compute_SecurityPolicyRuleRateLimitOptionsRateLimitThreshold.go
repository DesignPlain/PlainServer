package types

type Compute_SecurityPolicyRuleRateLimitOptionsRateLimitThreshold struct {
	// Number of HTTP(S) requests for calculating the threshold.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// Interval over which the threshold is computed.
	IntervalSec int `json:"intervalSec,omitempty" yaml:"intervalSec,omitempty"`
}
