package types

type Iot_TopicRuleErrorAction struct {
	//
	Dynamodb Iot_TopicRuleErrorActionDynamodb `json:"dynamodb,omitempty" yaml:"dynamodb,omitempty"`

	//
	Elasticsearch Iot_TopicRuleErrorActionElasticsearch `json:"elasticsearch,omitempty" yaml:"elasticsearch,omitempty"`

	//
	Http Iot_TopicRuleErrorActionHttp `json:"http,omitempty" yaml:"http,omitempty"`

	//
	Kinesis Iot_TopicRuleErrorActionKinesis `json:"kinesis,omitempty" yaml:"kinesis,omitempty"`

	//
	Republish Iot_TopicRuleErrorActionRepublish `json:"republish,omitempty" yaml:"republish,omitempty"`

	//
	StepFunctions Iot_TopicRuleErrorActionStepFunctions `json:"stepFunctions,omitempty" yaml:"stepFunctions,omitempty"`

	//
	CloudwatchAlarm Iot_TopicRuleErrorActionCloudwatchAlarm `json:"cloudwatchAlarm,omitempty" yaml:"cloudwatchAlarm,omitempty"`

	//
	Firehose Iot_TopicRuleErrorActionFirehose `json:"firehose,omitempty" yaml:"firehose,omitempty"`

	//
	Timestream Iot_TopicRuleErrorActionTimestream `json:"timestream,omitempty" yaml:"timestream,omitempty"`

	//
	CloudwatchLogs Iot_TopicRuleErrorActionCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	//
	Dynamodbv2 Iot_TopicRuleErrorActionDynamodbv2 `json:"dynamodbv2,omitempty" yaml:"dynamodbv2,omitempty"`

	//
	Lambda Iot_TopicRuleErrorActionLambda `json:"lambda,omitempty" yaml:"lambda,omitempty"`

	//
	Sqs Iot_TopicRuleErrorActionSqs `json:"sqs,omitempty" yaml:"sqs,omitempty"`

	//
	CloudwatchMetric Iot_TopicRuleErrorActionCloudwatchMetric `json:"cloudwatchMetric,omitempty" yaml:"cloudwatchMetric,omitempty"`

	//
	IotAnalytics Iot_TopicRuleErrorActionIotAnalytics `json:"iotAnalytics,omitempty" yaml:"iotAnalytics,omitempty"`

	//
	IotEvents Iot_TopicRuleErrorActionIotEvents `json:"iotEvents,omitempty" yaml:"iotEvents,omitempty"`

	//
	Kafka Iot_TopicRuleErrorActionKafka `json:"kafka,omitempty" yaml:"kafka,omitempty"`

	//
	S3 Iot_TopicRuleErrorActionS3 `json:"s3,omitempty" yaml:"s3,omitempty"`

	//
	Sns Iot_TopicRuleErrorActionSns `json:"sns,omitempty" yaml:"sns,omitempty"`
}
