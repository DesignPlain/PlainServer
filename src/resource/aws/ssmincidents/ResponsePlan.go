package ssmincidents

import types "DesignSphere_Server/src/resource/aws/types"

type ResponsePlan struct {
	// The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
	Engagements []string `json:"engagements,omitempty" yaml:"engagements,omitempty"`

	//
	IncidentTemplate types.Ssmincidents_ResponsePlanIncidentTemplate `json:"incidentTemplate,omitempty" yaml:"incidentTemplate,omitempty"`

	// Information about third-party services integrated into the response plan. The following values are supported:
	Integration types.Ssmincidents_ResponsePlanIntegration `json:"integration,omitempty" yaml:"integration,omitempty"`

	// The name of the response plan.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The tags applied to the response plan.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The actions that the response plan starts at the beginning of an incident.
	Action types.Ssmincidents_ResponsePlanAction `json:"action,omitempty" yaml:"action,omitempty"`

	// The Chatbot chat channel used for collaboration during an incident.
	ChatChannels []string `json:"chatChannels,omitempty" yaml:"chatChannels,omitempty"`

	// The long format of the response plan name. This field can contain spaces.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
