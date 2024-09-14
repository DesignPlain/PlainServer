package types

type Cloudfunctionsv2_FunctionEventTrigger struct {
	// Required. The type of event to observe.
	EventType string `json:"eventType,omitempty" yaml:"eventType,omitempty"`

	/*
	   The name of a Pub/Sub topic in the same project that will be used
	   as the transport topic for the event delivery.
	*/
	PubsubTopic string `json:"pubsubTopic,omitempty" yaml:"pubsubTopic,omitempty"`

	/*
	   Describes the retry policy in case of function's execution failure.
	   Retried execution is charged as any other execution.
	   Possible values are: `RETRY_POLICY_UNSPECIFIED`, `RETRY_POLICY_DO_NOT_RETRY`, `RETRY_POLICY_RETRY`.
	*/
	RetryPolicy string `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	/*
	   Optional. The email of the trigger's service account. The service account
	   must have permission to invoke Cloud Run services. If empty, defaults to the
	   Compute Engine default service account: {project_number}-compute@developer.gserviceaccount.com.
	*/
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	/*
	   (Output)
	   Output only. The resource name of the Eventarc trigger.
	*/
	Trigger string `json:"trigger,omitempty" yaml:"trigger,omitempty"`

	/*
	   The region that the trigger will be in. The trigger will only receive
	   events originating in this region. It can be the same
	   region as the function, a different region or multi-region, or the global
	   region. If not provided, defaults to the same region as the function.
	*/
	TriggerRegion string `json:"triggerRegion,omitempty" yaml:"triggerRegion,omitempty"`

	/*
	   Criteria used to filter events.
	   Structure is documented below.
	*/
	EventFilters []Cloudfunctionsv2_FunctionEventTriggerEventFilter `json:"eventFilters,omitempty" yaml:"eventFilters,omitempty"`
}
