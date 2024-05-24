package cloudwatch

import types "DesignSphere_Server/src/resource/aws/types"

type MetricAlarm struct {
	// The descriptive name for the alarm. This name must be unique within the user's AWS account
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The statistic to apply to the alarm's associated metric.
	   Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
	*/
	Statistic string `json:"statistic,omitempty" yaml:"statistic,omitempty"`

	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
	ActionsEnabled bool `json:"actionsEnabled,omitempty" yaml:"actionsEnabled,omitempty"`

	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic string `json:"extendedStatistic,omitempty" yaml:"extendedStatistic,omitempty"`

	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions []string `json:"insufficientDataActions,omitempty" yaml:"insufficientDataActions,omitempty"`

	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQueries []types.Cloudwatch_MetricAlarmMetricQuery `json:"metricQueries,omitempty" yaml:"metricQueries,omitempty"`

	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricId string `json:"thresholdMetricId,omitempty" yaml:"thresholdMetricId,omitempty"`

	// The description for the alarm.
	AlarmDescription string `json:"alarmDescription,omitempty" yaml:"alarmDescription,omitempty"`

	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm int `json:"datapointsToAlarm,omitempty" yaml:"datapointsToAlarm,omitempty"`

	/*
	   The namespace for the alarm's associated metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
	   See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	*/
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold float64 `json:"threshold,omitempty" yaml:"threshold,omitempty"`

	/*
	   The period in seconds over which the specified `statistic` is applied.
	   Valid values are `10`, `30`, or any multiple of `60`.
	*/
	Period int `json:"period,omitempty" yaml:"period,omitempty"`

	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
	ComparisonOperator string `json:"comparisonOperator,omitempty" yaml:"comparisonOperator,omitempty"`

	// The dimensions for the alarm's associated metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Dimensions map[string]string `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	/*
	   Used only for alarms based on percentiles.
	   If you specify `ignore`, the alarm state will not change during periods with too few data points to be statistically significant.
	   If you specify `evaluate` or omit this parameter, the alarm will always be evaluated and possibly change state no matter how many data points are available.
	   The following values are supported: `ignore`, and `evaluate`.
	*/
	EvaluateLowSampleCountPercentiles string `json:"evaluateLowSampleCountPercentiles,omitempty" yaml:"evaluateLowSampleCountPercentiles,omitempty"`

	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions []string `json:"okActions,omitempty" yaml:"okActions,omitempty"`

	// Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
	TreatMissingData string `json:"treatMissingData,omitempty" yaml:"treatMissingData,omitempty"`

	// The unit for the alarm's associated metric.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions []string `json:"alarmActions,omitempty" yaml:"alarmActions,omitempty"`

	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods int `json:"evaluationPeriods,omitempty" yaml:"evaluationPeriods,omitempty"`

	/*
	   The name for the alarm's associated metric.
	   See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	*/
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	/*
	   A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   > --NOTE:--  If you specify at least one `metric_query`, you may not specify a `metric_name`, `namespace`, `period` or `statistic`. If you do not specify a `metric_query`, you must specify each of these (although you may use `extended_statistic` instead of `statistic`).
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
