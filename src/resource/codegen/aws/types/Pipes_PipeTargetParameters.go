package types

type Pipes_PipeTargetParameters struct {
	// The parameters for using an CloudWatch Logs log stream as a target. Detailed below.
	CloudwatchLogsParameters Pipes_PipeTargetParametersCloudwatchLogsParameters `json:"cloudwatchLogsParameters,omitempty" yaml:"cloudwatchLogsParameters,omitempty"`

	// The parameters for using an EventBridge event bus as a target. Detailed below.
	EventbridgeEventBusParameters Pipes_PipeTargetParametersEventbridgeEventBusParameters `json:"eventbridgeEventBusParameters,omitempty" yaml:"eventbridgeEventBusParameters,omitempty"`

	// Valid JSON text passed to the target. In this case, nothing from the event itself is passed to the target. Maximum length of 8192 characters.
	InputTemplate string `json:"inputTemplate,omitempty" yaml:"inputTemplate,omitempty"`

	// The parameters for using a SageMaker pipeline as a target. Detailed below.
	SagemakerPipelineParameters Pipes_PipeTargetParametersSagemakerPipelineParameters `json:"sagemakerPipelineParameters,omitempty" yaml:"sagemakerPipelineParameters,omitempty"`

	// The parameters for using an AWS Batch job as a target. Detailed below.
	BatchJobParameters Pipes_PipeTargetParametersBatchJobParameters `json:"batchJobParameters,omitempty" yaml:"batchJobParameters,omitempty"`

	// These are custom parameter to be used when the target is an API Gateway REST APIs or EventBridge ApiDestinations. Detailed below.
	HttpParameters Pipes_PipeTargetParametersHttpParameters `json:"httpParameters,omitempty" yaml:"httpParameters,omitempty"`

	// The parameters for using a Kinesis stream as a source. Detailed below.
	KinesisStreamParameters Pipes_PipeTargetParametersKinesisStreamParameters `json:"kinesisStreamParameters,omitempty" yaml:"kinesisStreamParameters,omitempty"`

	// The parameters for using a Lambda function as a target. Detailed below.
	LambdaFunctionParameters Pipes_PipeTargetParametersLambdaFunctionParameters `json:"lambdaFunctionParameters,omitempty" yaml:"lambdaFunctionParameters,omitempty"`

	// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API BatchExecuteStatement. Detailed below.
	RedshiftDataParameters Pipes_PipeTargetParametersRedshiftDataParameters `json:"redshiftDataParameters,omitempty" yaml:"redshiftDataParameters,omitempty"`

	// The parameters for using a Amazon SQS stream as a target. Detailed below.
	SqsQueueParameters Pipes_PipeTargetParametersSqsQueueParameters `json:"sqsQueueParameters,omitempty" yaml:"sqsQueueParameters,omitempty"`

	// The parameters for using a Step Functions state machine as a target. Detailed below.
	StepFunctionStateMachineParameters Pipes_PipeTargetParametersStepFunctionStateMachineParameters `json:"stepFunctionStateMachineParameters,omitempty" yaml:"stepFunctionStateMachineParameters,omitempty"`

	// The parameters for using an Amazon ECS task as a target. Detailed below.
	EcsTaskParameters Pipes_PipeTargetParametersEcsTaskParameters `json:"ecsTaskParameters,omitempty" yaml:"ecsTaskParameters,omitempty"`
}
