package types

type Apigateway_UsagePlanThrottleSettings struct {
	//
	BurstLimit int `json:"burstLimit,omitempty" yaml:"burstLimit,omitempty"`

	//
	RateLimit float64 `json:"rateLimit,omitempty" yaml:"rateLimit,omitempty"`
}
