package types

type Dataform_RepositoryWorkflowConfigRecentScheduledExecutionRecord struct {
	/*
	   (Output)
	   The error status encountered upon this attempt to create the workflow invocation, if the attempt was unsuccessful.
	   Structure is documented below.
	*/
	ErrorStatuses []Dataform_RepositoryWorkflowConfigRecentScheduledExecutionRecordErrorStatus `json:"errorStatuses,omitempty" yaml:"errorStatuses,omitempty"`

	/*
	   (Output)
	   The timestamp of this workflow attempt.
	*/
	ExecutionTime string `json:"executionTime,omitempty" yaml:"executionTime,omitempty"`

	/*
	   (Output)
	   The name of the created workflow invocation, if one was successfully created. In the format projects/-/locations/-/repositories/-/workflowInvocations/-.
	*/
	WorkflowInvocation string `json:"workflowInvocation,omitempty" yaml:"workflowInvocation,omitempty"`
}
