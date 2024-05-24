package types

type Iot_TopicRuleStepFunction struct {
	// The name of the Step Functions state machine whose execution will be started.
	StateMachineName string `json:"stateMachineName,omitempty" yaml:"stateMachineName,omitempty"`

	// The prefix used to generate, along with a UUID, the unique state machine execution name.
	ExecutionNamePrefix string `json:"executionNamePrefix,omitempty" yaml:"executionNamePrefix,omitempty"`

	// The ARN of the IAM role that grants access to start execution of the state machine.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
