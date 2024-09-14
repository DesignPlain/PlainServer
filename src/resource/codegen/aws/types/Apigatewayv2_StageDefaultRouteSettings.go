package types

type Apigatewayv2_StageDefaultRouteSettings struct {
	// Throttling rate limit for the default route.
	ThrottlingRateLimit float64 `json:"throttlingRateLimit,omitempty" yaml:"throttlingRateLimit,omitempty"`

	/*
	   Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	   Defaults to `false`. Supported only for WebSocket APIs.
	*/
	DataTraceEnabled bool `json:"dataTraceEnabled,omitempty" yaml:"dataTraceEnabled,omitempty"`

	// Whether detailed metrics are enabled for the default route. Defaults to `false`.
	DetailedMetricsEnabled bool `json:"detailedMetricsEnabled,omitempty" yaml:"detailedMetricsEnabled,omitempty"`

	/*
	   Logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	   Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
	*/
	LoggingLevel string `json:"loggingLevel,omitempty" yaml:"loggingLevel,omitempty"`

	// Throttling burst limit for the default route.
	ThrottlingBurstLimit int `json:"throttlingBurstLimit,omitempty" yaml:"throttlingBurstLimit,omitempty"`
}
