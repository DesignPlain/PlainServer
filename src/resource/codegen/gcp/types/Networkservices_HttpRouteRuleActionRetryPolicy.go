package types

type Networkservices_HttpRouteRuleActionRetryPolicy struct {
	// Specifies a non-zero timeout per retry attempt. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	PerTryTimeout string `json:"perTryTimeout,omitempty" yaml:"perTryTimeout,omitempty"`

	// Specifies one or more conditions when this retry policy applies.
	RetryConditions []string `json:"retryConditions,omitempty" yaml:"retryConditions,omitempty"`

	// Specifies the allowed number of retries.
	NumRetries int `json:"numRetries,omitempty" yaml:"numRetries,omitempty"`
}
