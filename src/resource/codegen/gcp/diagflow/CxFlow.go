package diagflow

import types "libds/gcp/types"

type CxFlow struct {
	/*
	   Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent. When you create an agent, the Default Start Flow is created automatically.
	   The Default Start Flow cannot be deleted; deleting the `gcp.diagflow.CxFlow` resource does nothing to the underlying GCP resources.

	   > Avoid having multiple `gcp.diagflow.CxFlow` resources linked to the same agent with `is_default_start_flow = true` because they will compete to control a single Default Start Flow resource in GCP.
	*/
	IsDefaultStartFlow bool `json:"isDefaultStartFlow,omitempty" yaml:"isDefaultStartFlow,omitempty"`

	/*
	   Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at the higher level.
	   Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	   Structure is documented below.
	*/
	AdvancedSettings types.Diagflow_CxFlowAdvancedSettings `json:"advancedSettings,omitempty" yaml:"advancedSettings,omitempty"`

	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The language of the following fields in flow:
	   Flow.event_handlers.trigger_fulfillment.messages
	   Flow.event_handlers.trigger_fulfillment.conditional_cases
	   Flow.transition_routes.trigger_fulfillment.messages
	   Flow.transition_routes.trigger_fulfillment.conditional_cases
	   If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	*/
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	/*
	   NLU related settings of the flow.
	   Structure is documented below.
	*/
	NluSettings types.Diagflow_CxFlowNluSettings `json:"nluSettings,omitempty" yaml:"nluSettings,omitempty"`

	/*
	   The agent to create a flow for.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   A flow's transition route group serve two purposes:
	   They are responsible for matching the user's first utterances in the flow.
	   They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
	   Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	*/
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty" yaml:"transitionRouteGroups,omitempty"`

	/*
	   A flow's transition routes serve two purposes:
	   They are responsible for matching the user's first utterances in the flow.
	   They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
	   TransitionRoutes are evalauted in the following order:
	   TransitionRoutes with intent specified.
	   TransitionRoutes with only condition specified.
	   TransitionRoutes with intent specified are inherited by pages in the flow.
	   Structure is documented below.
	*/
	TransitionRoutes []types.Diagflow_CxFlowTransitionRoute `json:"transitionRoutes,omitempty" yaml:"transitionRoutes,omitempty"`

	/*
	   The human-readable name of the flow.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   A flow's event handlers serve two purposes:
	   They are responsible for handling events (e.g. no match, webhook errors) in the flow.
	   They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
	   Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	   Structure is documented below.
	*/
	EventHandlers []types.Diagflow_CxFlowEventHandler `json:"eventHandlers,omitempty" yaml:"eventHandlers,omitempty"`
}
