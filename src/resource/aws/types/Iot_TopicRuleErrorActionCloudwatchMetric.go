package types

type Iot_TopicRuleErrorActionCloudwatchMetric struct {
	// The CloudWatch metric name.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// The CloudWatch metric namespace name.
	MetricNamespace string `json:"metricNamespace,omitempty" yaml:"metricNamespace,omitempty"`

	// An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
	MetricTimestamp string `json:"metricTimestamp,omitempty" yaml:"metricTimestamp,omitempty"`

	// The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
	MetricUnit string `json:"metricUnit,omitempty" yaml:"metricUnit,omitempty"`

	// The CloudWatch metric value.
	MetricValue string `json:"metricValue,omitempty" yaml:"metricValue,omitempty"`

	// The IAM role ARN that allows access to the CloudWatch metric.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
