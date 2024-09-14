package types

type Diagflow_CxFlowTransitionRoute struct {
	/*
	   The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.
	   Structure is documented below.
	*/
	TriggerFulfillment Diagflow_CxFlowTransitionRouteTriggerFulfillment `json:"triggerFulfillment,omitempty" yaml:"triggerFulfillment,omitempty"`

	/*
	   The condition to evaluate against form parameters or session parameters.
	   At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	*/
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`

	/*
	   The unique identifier of an Intent.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	*/
	Intent string `json:"intent,omitempty" yaml:"intent,omitempty"`

	/*
	   (Output)
	   The unique identifier of this transition route.
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
}
