package types

type Apigatewayv2_StageRouteSetting struct {
	/*
	   Logging level for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	   Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
	*/
	LoggingLevel string `json:"loggingLevel,omitempty" yaml:"loggingLevel,omitempty"`

	// Route key.
	RouteKey string `json:"routeKey,omitempty" yaml:"routeKey,omitempty"`

	// Throttling burst limit for the route.
	ThrottlingBurstLimit int `json:"throttlingBurstLimit,omitempty" yaml:"throttlingBurstLimit,omitempty"`

	// Throttling rate limit for the route.
	ThrottlingRateLimit float64 `json:"throttlingRateLimit,omitempty" yaml:"throttlingRateLimit,omitempty"`

	/*
	   Whether data trace logging is enabled for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	   Defaults to `false`. Supported only for WebSocket APIs.
	*/
	DataTraceEnabled bool `json:"dataTraceEnabled,omitempty" yaml:"dataTraceEnabled,omitempty"`

	// Whether detailed metrics are enabled for the route. Defaults to `false`.
	DetailedMetricsEnabled bool `json:"detailedMetricsEnabled,omitempty" yaml:"detailedMetricsEnabled,omitempty"`
}
