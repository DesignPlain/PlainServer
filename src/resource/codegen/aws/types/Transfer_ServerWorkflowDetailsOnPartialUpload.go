package types

type Transfer_ServerWorkflowDetailsOnPartialUpload struct {
	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	ExecutionRole string `json:"executionRole,omitempty" yaml:"executionRole,omitempty"`

	// A unique identifier for the workflow.
	WorkflowId string `json:"workflowId,omitempty" yaml:"workflowId,omitempty"`
}
