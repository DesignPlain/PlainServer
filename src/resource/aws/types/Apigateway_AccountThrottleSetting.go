package types

type Apigateway_AccountThrottleSetting struct {
	// Absolute maximum number of times API Gateway allows the API to be called per second (RPS).
	BurstLimit int `json:"burstLimit,omitempty" yaml:"burstLimit,omitempty"`

	// Number of times API Gateway allows the API to be called per second on average (RPS).
	RateLimit float64 `json:"rateLimit,omitempty" yaml:"rateLimit,omitempty"`
}
