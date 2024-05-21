package types

type Cloudfunctions_getFunctionEventTrigger struct {
	/*
	   The type of event to observe. For example: `"google.storage.object.finalize"`.
	   See the documentation on [calling Cloud Functions](https://cloud.google.com/functions/docs/calling/)
	   for a full reference of accepted triggers.
	*/
	EventType string `json:"eventType,omitempty" yaml:"eventType,omitempty"`

	// Policy for failed executions. Structure is documented below.
	FailurePolicies []Cloudfunctions_getFunctionEventTriggerFailurePolicy `json:"failurePolicies,omitempty" yaml:"failurePolicies,omitempty"`

	// The name of the resource whose events are being observed, for example, `"myBucket"`
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
