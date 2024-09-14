package types

type Sfn_StateMachineLoggingConfiguration struct {
	// Determines whether execution data is included in your log. When set to `false`, data is excluded.
	IncludeExecutionData bool `json:"includeExecutionData,omitempty" yaml:"includeExecutionData,omitempty"`

	// Defines which category of execution history events are logged. Valid values: `ALL`, `ERROR`, `FATAL`, `OFF`
	Level string `json:"level,omitempty" yaml:"level,omitempty"`

	// Amazon Resource Name (ARN) of a CloudWatch log group. Make sure the State Machine has the correct IAM policies for logging. The ARN must end with `:-`
	LogDestination string `json:"logDestination,omitempty" yaml:"logDestination,omitempty"`
}
