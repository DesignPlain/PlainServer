package types

type Transfer_ServerWorkflowDetails struct {
	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded. See `on_upload` Block below for details.
	OnUpload Transfer_ServerWorkflowDetailsOnUpload `json:"onUpload,omitempty" yaml:"onUpload,omitempty"`

	// A trigger that starts a workflow if a file is only partially uploaded. See Workflow Detail below. See `on_partial_upload` Block below for details.
	OnPartialUpload Transfer_ServerWorkflowDetailsOnPartialUpload `json:"onPartialUpload,omitempty" yaml:"onPartialUpload,omitempty"`
}
