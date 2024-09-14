package transfer

import types "libds/aws/types"

type Workflow struct {
	// Specifies the steps (actions) to take if errors are encountered during execution of the workflow. See Workflow Steps below.
	OnExceptionSteps []types.Transfer_WorkflowOnExceptionStep `json:"onExceptionSteps,omitempty" yaml:"onExceptionSteps,omitempty"`

	// Specifies the details for the steps that are in the specified workflow. See Workflow Steps below.
	Steps []types.Transfer_WorkflowStep `json:"steps,omitempty" yaml:"steps,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A textual description for the workflow.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
