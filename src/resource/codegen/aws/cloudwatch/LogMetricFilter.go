package cloudwatch

import types "libds/aws/types"

type LogMetricFilter struct {
	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation types.Cloudwatch_LogMetricFilterMetricTransformation `json:"metricTransformation,omitempty" yaml:"metricTransformation,omitempty"`

	// A name for the metric filter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	   for extracting metric data out of ingested log events.
	*/
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// The name of the log group to associate the metric filter with.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`
}
