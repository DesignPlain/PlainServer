package types

type Fis_ExperimentTemplateAction struct {
	// Action's target, if applicable. See below.
	Target Fis_ExperimentTemplateActionTarget `json:"target,omitempty" yaml:"target,omitempty"`

	// ID of the action. To find out what actions are supported see [AWS FIS actions reference](https://docs.aws.amazon.com/fis/latest/userguide/fis-actions-reference.html).
	ActionId string `json:"actionId,omitempty" yaml:"actionId,omitempty"`

	// Description of the action.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Friendly name of the action.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Parameter(s) for the action, if applicable. See below.
	Parameters []Fis_ExperimentTemplateActionParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Set of action names that must complete before this action can be executed.
	StartAfters []string `json:"startAfters,omitempty" yaml:"startAfters,omitempty"`
}
