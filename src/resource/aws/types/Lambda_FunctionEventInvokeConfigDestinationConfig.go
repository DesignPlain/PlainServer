package types

type Lambda_FunctionEventInvokeConfigDestinationConfig struct {
	// Configuration block with destination configuration for failed asynchronous invocations. See below for details.
	OnFailure Lambda_FunctionEventInvokeConfigDestinationConfigOnFailure `json:"onFailure,omitempty" yaml:"onFailure,omitempty"`

	// Configuration block with destination configuration for successful asynchronous invocations. See below for details.
	OnSuccess Lambda_FunctionEventInvokeConfigDestinationConfigOnSuccess `json:"onSuccess,omitempty" yaml:"onSuccess,omitempty"`
}
