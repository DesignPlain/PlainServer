package types

type Identityplatform_ConfigBlockingFunctionsTrigger struct {
	// The identifier for this object. Format specified above.
	EventType string `json:"eventType,omitempty" yaml:"eventType,omitempty"`

	// HTTP URI trigger for the Cloud Function.
	FunctionUri string `json:"functionUri,omitempty" yaml:"functionUri,omitempty"`

	/*
	   (Output)
	   When the trigger was changed.
	*/
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
