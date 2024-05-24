package cloudwatch

import types "DesignSphere_Server/src/resource/aws/types"

type EventTarget struct {
	// Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
	RetryPolicy types.Cloudwatch_EventTargetRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	/*
	   The name of the rule you want to add targets to.

	   The following arguments are optional:
	*/
	Rule string `json:"rule,omitempty" yaml:"rule,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget types.Cloudwatch_EventTargetSqsTarget `json:"sqsTarget,omitempty" yaml:"sqsTarget,omitempty"`

	// The unique target assignment ID. If missing, will generate a random, unique id.
	TargetId string `json:"targetId,omitempty" yaml:"targetId,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget types.Cloudwatch_EventTargetBatchTarget `json:"batchTarget,omitempty" yaml:"batchTarget,omitempty"`

	// Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
	DeadLetterConfig types.Cloudwatch_EventTargetDeadLetterConfig `json:"deadLetterConfig,omitempty" yaml:"deadLetterConfig,omitempty"`

	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget types.Cloudwatch_EventTargetEcsTarget `json:"ecsTarget,omitempty" yaml:"ecsTarget,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget types.Cloudwatch_EventTargetKinesisTarget `json:"kinesisTarget,omitempty" yaml:"kinesisTarget,omitempty"`

	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `input_transformer`.
	InputPath string `json:"inputPath,omitempty" yaml:"inputPath,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used or target in `arn` is EC2 instance, Kinesis data stream, Step Functions state machine, or Event Bus in different account or region.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon SageMaker Pipeline. Documented below. A maximum of 1 are allowed.
	SagemakerPipelineTarget types.Cloudwatch_EventTargetSagemakerPipelineTarget `json:"sagemakerPipelineTarget,omitempty" yaml:"sagemakerPipelineTarget,omitempty"`

	/*
	   The name or ARN of the event bus to associate with the rule.
	   If you omit this, the `default` event bus is used.
	*/
	EventBusName string `json:"eventBusName,omitempty" yaml:"eventBusName,omitempty"`

	// Valid JSON text passed to the target. Conflicts with `input_path` and `input_transformer`.
	Input string `json:"input,omitempty" yaml:"input,omitempty"`

	// Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `input_path`.
	InputTransformer types.Cloudwatch_EventTargetInputTransformer `json:"inputTransformer,omitempty" yaml:"inputTransformer,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Redshift Statement. Documented below. A maximum of 1 are allowed.
	RedshiftTarget types.Cloudwatch_EventTargetRedshiftTarget `json:"redshiftTarget,omitempty" yaml:"redshiftTarget,omitempty"`

	// The Amazon Resource Name (ARN) of the target.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
	HttpTarget types.Cloudwatch_EventTargetHttpTarget `json:"httpTarget,omitempty" yaml:"httpTarget,omitempty"`

	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets []types.Cloudwatch_EventTargetRunCommandTarget `json:"runCommandTargets,omitempty" yaml:"runCommandTargets,omitempty"`
}
