package types

type Diagflow_CxFlowTransitionRouteTriggerFulfillmentSetParameterAction struct {
	// Display name of the parameter.
	Parameter string `json:"parameter,omitempty" yaml:"parameter,omitempty"`

	// The new JSON-encoded value of the parameter. A null value clears the parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
