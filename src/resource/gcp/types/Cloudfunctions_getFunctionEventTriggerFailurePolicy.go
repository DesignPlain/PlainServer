package types

type Cloudfunctions_getFunctionEventTriggerFailurePolicy struct {
	// Whether the function should be retried on failure.
	Retry bool `json:"retry,omitempty" yaml:"retry,omitempty"`
}
