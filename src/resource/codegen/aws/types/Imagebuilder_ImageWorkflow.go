package types

type Imagebuilder_ImageWorkflow struct {
	// Configuration block for the workflow parameters. Detailed below.
	Parameters []Imagebuilder_ImageWorkflowParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the Image Builder Workflow.

	   The following arguments are optional:
	*/
	WorkflowArn string `json:"workflowArn,omitempty" yaml:"workflowArn,omitempty"`

	// The action to take if the workflow fails. Must be one of `CONTINUE` or `ABORT`.
	OnFailure string `json:"onFailure,omitempty" yaml:"onFailure,omitempty"`

	// The parallel group in which to run a test Workflow.
	ParallelGroup string `json:"parallelGroup,omitempty" yaml:"parallelGroup,omitempty"`
}
