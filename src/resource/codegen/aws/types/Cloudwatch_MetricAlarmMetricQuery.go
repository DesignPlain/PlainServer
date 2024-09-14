package types

type Cloudwatch_MetricAlarmMetricQuery struct {
	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the id of the other metrics to refer to those metrics, and can also use the id of other expressions to use the result of those expressions. For more information about metric math expressions, see Metric Math Syntax and Functions in the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax).
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// A short name used to tie this object to the results in the response. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`

	// The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
	Metric Cloudwatch_MetricAlarmMetricQueryMetric `json:"metric,omitempty" yaml:"metric,omitempty"`

	/*
	   Granularity in seconds of returned data points.
	   For metrics with regular resolution, valid values are any multiple of `60`.
	   For high-resolution metrics, valid values are `1`, `5`, `10`, `30`, or any multiple of `60`.
	*/
	Period int `json:"period,omitempty" yaml:"period,omitempty"`

	/*
	   Specify exactly one `metric_query` to be `true` to use that `metric_query` result as the alarm.

	   > --NOTE:--  You must specify either `metric` or `expression`. Not both.
	*/
	ReturnData bool `json:"returnData,omitempty" yaml:"returnData,omitempty"`
}
