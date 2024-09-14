package types

type Networkservices_GrpcRouteRuleActionRetryPolicy struct {
	/*
	   Specifies one or more conditions when this retry policy applies.
	   Each value may be one of: `connect-failure`, `refused-stream`, `cancelled`, `deadline-exceeded`, `resource-exhausted`, `unavailable`.
	*/
	RetryConditions []string `json:"retryConditions,omitempty" yaml:"retryConditions,omitempty"`

	/*
	   Specifies the allowed number of retries.

	   - - -
	*/
	NumRetries int `json:"numRetries,omitempty" yaml:"numRetries,omitempty"`
}
