package types

type Scheduler_ScheduleTarget struct {
	// Templated target type for the Amazon ECS [`RunTask`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) API operation. Detailed below.
	EcsParameters Scheduler_ScheduleTargetEcsParameters `json:"ecsParameters,omitempty" yaml:"ecsParameters,omitempty"`

	// Templated target type for the EventBridge [`PutEvents`](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html) API operation. Detailed below.
	EventbridgeParameters Scheduler_ScheduleTargetEventbridgeParameters `json:"eventbridgeParameters,omitempty" yaml:"eventbridgeParameters,omitempty"`

	// Templated target type for the Amazon Kinesis [`PutRecord`](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html) API operation. Detailed below.
	KinesisParameters Scheduler_ScheduleTargetKinesisParameters `json:"kinesisParameters,omitempty" yaml:"kinesisParameters,omitempty"`

	/*
	   ARN of the IAM role that EventBridge Scheduler will use for this target when the schedule is invoked. Read more in [Set up the execution role](https://docs.aws.amazon.com/scheduler/latest/UserGuide/setting-up.html#setting-up-execution-role).

	   The following arguments are optional:
	*/
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Templated target type for the Amazon SageMaker [`StartPipelineExecution`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StartPipelineExecution.html) API operation. Detailed below.
	SagemakerPipelineParameters Scheduler_ScheduleTargetSagemakerPipelineParameters `json:"sagemakerPipelineParameters,omitempty" yaml:"sagemakerPipelineParameters,omitempty"`

	// Information about an Amazon SQS queue that EventBridge Scheduler uses as a dead-letter queue for your schedule. If specified, EventBridge Scheduler delivers failed events that could not be successfully delivered to a target to the queue. Detailed below.
	DeadLetterConfig Scheduler_ScheduleTargetDeadLetterConfig `json:"deadLetterConfig,omitempty" yaml:"deadLetterConfig,omitempty"`

	// Text, or well-formed JSON, passed to the target. Read more in [Universal target](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets-universal.html).
	Input string `json:"input,omitempty" yaml:"input,omitempty"`

	// Information about the retry policy settings. Detailed below.
	RetryPolicy Scheduler_ScheduleTargetRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	// The templated target type for the Amazon SQS [`SendMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html) API operation. Detailed below.
	SqsParameters Scheduler_ScheduleTargetSqsParameters `json:"sqsParameters,omitempty" yaml:"sqsParameters,omitempty"`

	// ARN of the target of this schedule, such as a SQS queue or ECS cluster. For universal targets, this is a [Service ARN specific to the target service](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets-universal.html#supported-universal-targets).
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`
}
