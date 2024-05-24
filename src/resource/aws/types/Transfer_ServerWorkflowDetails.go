package types

type Transfer_ServerWorkflowDetails struct {
	// A trigger that starts a workflow if a file is only partially uploaded. See Workflow Detail below.
	OnPartialUpload Transfer_ServerWorkflowDetailsOnPartialUpload `json:"onPartialUpload,omitempty" yaml:"onPartialUpload,omitempty"`

	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded. See Workflow Detail below.
	OnUpload Transfer_ServerWorkflowDetailsOnUpload `json:"onUpload,omitempty" yaml:"onUpload,omitempty"`
}
