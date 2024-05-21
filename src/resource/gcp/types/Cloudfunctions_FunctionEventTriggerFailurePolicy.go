package types

type Cloudfunctions_FunctionEventTriggerFailurePolicy struct {
	// Whether the function should be retried on failure. Defaults to `false`.
	Retry bool `json:"retry,omitempty" yaml:"retry,omitempty"`
}
