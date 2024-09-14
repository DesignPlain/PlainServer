package route53

type HealthCheck struct {
	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval int `json:"requestInterval,omitempty" yaml:"requestInterval,omitempty"`

	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath string `json:"resourcePath,omitempty" yaml:"resourcePath,omitempty"`

	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
	SearchString string `json:"searchString,omitempty" yaml:"searchString,omitempty"`

	// The name of the CloudWatch alarm.
	CloudwatchAlarmName string `json:"cloudwatchAlarmName,omitempty" yaml:"cloudwatchAlarmName,omitempty"`

	/*
	   A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	   - For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
	   - For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
	   - For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.

	   > --Note:-- After you disable a health check, Route 53 considers the status of the health check to always be healthy. If you configured DNS failover, Route 53 continues to route traffic to the corresponding resources. If you want to stop routing traffic to a resource, change the value of `invert_healthcheck`.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// The fully qualified domain name of the endpoint to be checked. If a value is set for `ip_address`, the value set for `fqdn` will be passed in the `Host` header.
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	// The IP address of the endpoint to be checked.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`

	// The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS' defaults: when the `type` is "HTTPS" `enable_sni` defaults to `true`, when `type` is anything else `enable_sni` defaults to `false`.
	EnableSni bool `json:"enableSni,omitempty" yaml:"enableSni,omitempty"`

	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold int `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`

	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency bool `json:"measureLatency,omitempty" yaml:"measureLatency,omitempty"`

	// The port of the endpoint to be checked.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// A map of tags to assign to the health check. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   This is a reference name used in Caller Reference
	   (helpful for identifying single health_check set amongst others)
	*/
	ReferenceName string `json:"referenceName,omitempty" yaml:"referenceName,omitempty"`

	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
	RoutingControlArn string `json:"routingControlArn,omitempty" yaml:"routingControlArn,omitempty"`

	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold int `json:"childHealthThreshold,omitempty" yaml:"childHealthThreshold,omitempty"`

	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks []string `json:"childHealthchecks,omitempty" yaml:"childHealthchecks,omitempty"`

	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion string `json:"cloudwatchAlarmRegion,omitempty" yaml:"cloudwatchAlarmRegion,omitempty"`

	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
	InsufficientDataHealthStatus string `json:"insufficientDataHealthStatus,omitempty" yaml:"insufficientDataHealthStatus,omitempty"`

	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck bool `json:"invertHealthcheck,omitempty" yaml:"invertHealthcheck,omitempty"`
}
