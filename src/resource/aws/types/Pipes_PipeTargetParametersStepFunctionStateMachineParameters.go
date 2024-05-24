package types

type Pipes_PipeTargetParametersStepFunctionStateMachineParameters struct {
	// Specify whether to invoke the function synchronously or asynchronously. Valid Values: REQUEST_RESPONSE, FIRE_AND_FORGET.
	InvocationType string `json:"invocationType,omitempty" yaml:"invocationType,omitempty"`
}
