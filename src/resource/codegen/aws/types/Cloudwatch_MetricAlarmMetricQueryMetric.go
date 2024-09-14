package types

type Cloudwatch_MetricAlarmMetricQueryMetric struct {
	/*
	   The statistic to apply to this metric.
	   See docs for [supported statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html).
	*/
	Stat string `json:"stat,omitempty" yaml:"stat,omitempty"`

	// The unit for this metric.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Dimensions map[string]string `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	/*
	   The name for this metric.
	   See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	*/
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	/*
	   The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
	   See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	*/
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	/*
	   Granularity in seconds of returned data points.
	   For metrics with regular resolution, valid values are any multiple of `60`.
	   For high-resolution metrics, valid values are `1`, `5`, `10`, `30`, or any multiple of `60`.
	*/
	Period int `json:"period,omitempty" yaml:"period,omitempty"`
}
