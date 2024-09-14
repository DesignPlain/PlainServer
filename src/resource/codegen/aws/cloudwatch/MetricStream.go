package cloudwatch

import types "libds/aws/types"

type MetricStream struct {
	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `include_filter`.
	ExcludeFilters []types.Cloudwatch_MetricStreamExcludeFilter `json:"excludeFilters,omitempty" yaml:"excludeFilters,omitempty"`

	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	FirehoseArn string `json:"firehoseArn,omitempty" yaml:"firehoseArn,omitempty"`

	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
	IncludeLinkedAccountsMetrics bool `json:"includeLinkedAccountsMetrics,omitempty" yaml:"includeLinkedAccountsMetrics,omitempty"`

	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   Output format for the stream. Possible values are `json`, `opentelemetry0.7`, and `opentelemetry1.0`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).

	   The following arguments are optional:
	*/
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`

	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `output_format`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7` or `opentelemetry1.0`, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfigurations []types.Cloudwatch_MetricStreamStatisticsConfiguration `json:"statisticsConfigurations,omitempty" yaml:"statisticsConfigurations,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `exclude_filter`.
	IncludeFilters []types.Cloudwatch_MetricStreamIncludeFilter `json:"includeFilters,omitempty" yaml:"includeFilters,omitempty"`

	// Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
