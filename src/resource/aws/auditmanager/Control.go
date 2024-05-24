package auditmanager

import types "DesignSphere_Server/src/resource/aws/types"

type Control struct {
	// Title of the action plan for remediating the control.
	ActionPlanTitle string `json:"actionPlanTitle,omitempty" yaml:"actionPlanTitle,omitempty"`

	/*
	   Data mapping sources. See `control_mapping_sources` below.

	   The following arguments are optional:
	*/
	ControlMappingSources []types.Auditmanager_ControlControlMappingSource `json:"controlMappingSources,omitempty" yaml:"controlMappingSources,omitempty"`

	// Description of the control.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the control.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Steps to follow to determine if the control is satisfied.
	TestingInformation string `json:"testingInformation,omitempty" yaml:"testingInformation,omitempty"`

	// Recommended actions to carry out if the control isn't fulfilled.
	ActionPlanInstructions string `json:"actionPlanInstructions,omitempty" yaml:"actionPlanInstructions,omitempty"`
}
