package ec2

type VpcNetworkPerformanceMetricSubscription struct {
	// The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
	Statistic string `json:"statistic,omitempty" yaml:"statistic,omitempty"`

	// The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
	Metric string `json:"metric,omitempty" yaml:"metric,omitempty"`

	// The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
