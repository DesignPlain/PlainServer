package types

type Appengine_FlexibleAppVersionAutomaticScalingRequestUtilization struct {
	// Target number of concurrent requests.
	TargetConcurrentRequests float64 `json:"targetConcurrentRequests,omitempty" yaml:"targetConcurrentRequests,omitempty"`

	// Target requests per second.
	TargetRequestCountPerSecond string `json:"targetRequestCountPerSecond,omitempty" yaml:"targetRequestCountPerSecond,omitempty"`
}
