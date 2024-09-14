package ssmincidents

import types "libds/aws/types"

type ResponsePlan struct {
	//
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	//
	Engagements []string `json:"engagements,omitempty" yaml:"engagements,omitempty"`

	//
	IncidentTemplate types.Ssmincidents_ResponsePlanIncidentTemplate `json:"incidentTemplate,omitempty" yaml:"incidentTemplate,omitempty"`

	//
	Integration types.Ssmincidents_ResponsePlanIntegration `json:"integration,omitempty" yaml:"integration,omitempty"`

	// The name of the response plan.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Action types.Ssmincidents_ResponsePlanAction `json:"action,omitempty" yaml:"action,omitempty"`

	//
	ChatChannels []string `json:"chatChannels,omitempty" yaml:"chatChannels,omitempty"`
}
