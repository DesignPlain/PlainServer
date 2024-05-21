package types

type Diagflow_CxPageEventHandler struct {
	// The name of the event to handle.
	Event string `json:"event,omitempty" yaml:"event,omitempty"`

	/*
	   (Output)
	   The unique identifier of this event handler.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The target flow to transition to.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	*/
	TargetFlow string `json:"targetFlow,omitempty" yaml:"targetFlow,omitempty"`

	/*
	   The target page to transition to.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	*/
	TargetPage string `json:"targetPage,omitempty" yaml:"targetPage,omitempty"`

	/*
	   The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	   Structure is documented below.
	*/
	TriggerFulfillment Diagflow_CxPageEventHandlerTriggerFulfillment `json:"triggerFulfillment,omitempty" yaml:"triggerFulfillment,omitempty"`
}
