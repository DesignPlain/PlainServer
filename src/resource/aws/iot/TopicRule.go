package iot

import types "DesignSphere_Server/src/resource/aws/types"

type TopicRule struct {
	// Specifies whether the rule is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Configuration block with error action to be associated with the rule. See the documentation for `cloudwatch_alarm`, `cloudwatch_logs`, `cloudwatch_metric`, `dynamodb`, `dynamodbv2`, `elasticsearch`, `firehose`, `http`, `iot_analytics`, `iot_events`, `kafka`, `kinesis`, `lambda`, `republish`, `s3`, `sns`, `sqs`, `step_functions`, `timestream` configuration blocks for further configuration details.
	ErrorAction types.Iot_TopicRuleErrorAction `json:"errorAction,omitempty" yaml:"errorAction,omitempty"`

	//
	Https []types.Iot_TopicRuleHttp `json:"https,omitempty" yaml:"https,omitempty"`

	//
	Kineses []types.Iot_TopicRuleKinesis `json:"kineses,omitempty" yaml:"kineses,omitempty"`

	//
	Republishes []types.Iot_TopicRuleRepublish `json:"republishes,omitempty" yaml:"republishes,omitempty"`

	//
	Dynamodbs []types.Iot_TopicRuleDynamodb `json:"dynamodbs,omitempty" yaml:"dynamodbs,omitempty"`

	//
	Dynamodbv2s []types.Iot_TopicRuleDynamodbv2 `json:"dynamodbv2s,omitempty" yaml:"dynamodbv2s,omitempty"`

	//
	IotAnalytics []types.Iot_TopicRuleIotAnalytic `json:"iotAnalytics,omitempty" yaml:"iotAnalytics,omitempty"`

	//
	IotEvents []types.Iot_TopicRuleIotEvent `json:"iotEvents,omitempty" yaml:"iotEvents,omitempty"`

	// The name of the rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The version of the SQL rules engine to use when evaluating the rule.
	SqlVersion string `json:"sqlVersion,omitempty" yaml:"sqlVersion,omitempty"`

	//
	Sqs []types.Iot_TopicRuleSqs `json:"sqs,omitempty" yaml:"sqs,omitempty"`

	//
	StepFunctions []types.Iot_TopicRuleStepFunction `json:"stepFunctions,omitempty" yaml:"stepFunctions,omitempty"`

	//
	CloudwatchAlarms []types.Iot_TopicRuleCloudwatchAlarm `json:"cloudwatchAlarms,omitempty" yaml:"cloudwatchAlarms,omitempty"`

	//
	Elasticsearch []types.Iot_TopicRuleElasticsearch `json:"elasticsearch,omitempty" yaml:"elasticsearch,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Lambdas []types.Iot_TopicRuleLambda `json:"lambdas,omitempty" yaml:"lambdas,omitempty"`

	//
	S3 []types.Iot_TopicRuleS3 `json:"s3,omitempty" yaml:"s3,omitempty"`

	//
	Sns []types.Iot_TopicRuleSns `json:"sns,omitempty" yaml:"sns,omitempty"`

	//
	Timestreams []types.Iot_TopicRuleTimestream `json:"timestreams,omitempty" yaml:"timestreams,omitempty"`

	//
	CloudwatchLogs []types.Iot_TopicRuleCloudwatchLog `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	// The description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	Kafkas []types.Iot_TopicRuleKafka `json:"kafkas,omitempty" yaml:"kafkas,omitempty"`

	// The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
	Sql string `json:"sql,omitempty" yaml:"sql,omitempty"`

	//
	CloudwatchMetrics []types.Iot_TopicRuleCloudwatchMetric `json:"cloudwatchMetrics,omitempty" yaml:"cloudwatchMetrics,omitempty"`

	//
	Firehoses []types.Iot_TopicRuleFirehose `json:"firehoses,omitempty" yaml:"firehoses,omitempty"`
}
