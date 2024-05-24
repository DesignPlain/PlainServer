package types

type Apigateway_UsagePlanThrottleSettings struct {
	// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
	BurstLimit int `json:"burstLimit,omitempty" yaml:"burstLimit,omitempty"`

	// The API request steady-state rate limit.
	RateLimit float64 `json:"rateLimit,omitempty" yaml:"rateLimit,omitempty"`
}
