package types

type Wafv2_RuleGroupRuleVisibilityConfig struct {
	// A boolean indicating whether the associated resource sends metrics to CloudWatch. For the list of available metrics, see [AWS WAF Metrics](https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html#waf-metrics).
	CloudwatchMetricsEnabled bool `json:"cloudwatchMetricsEnabled,omitempty" yaml:"cloudwatchMetricsEnabled,omitempty"`

	// A friendly name of the CloudWatch metric. The name can contain only alphanumeric characters (A-Z, a-z, 0-9) hyphen(-) and underscore (_), with length from one to 128 characters. It can't contain whitespace or metric names reserved for AWS WAF, for example `All` and `Default_Action`.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules. You can view the sampled requests through the AWS WAF console.
	SampledRequestsEnabled bool `json:"sampledRequestsEnabled,omitempty" yaml:"sampledRequestsEnabled,omitempty"`
}
