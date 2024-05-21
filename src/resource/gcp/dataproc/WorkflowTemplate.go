package dataproc

import types "DesignSphere_Server/src/resource/gcp/types"

type WorkflowTemplate struct {
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs []types.Dataproc_WorkflowTemplateJob `json:"jobs,omitempty" yaml:"jobs,omitempty"`

	// Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
	DagTimeout string `json:"dagTimeout,omitempty" yaml:"dagTimeout,omitempty"`

	// The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance. Label --keys-- must contain 1 to 63 characters, and must conform to (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a template.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Output only. The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. - For `projects.regions.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` - For `projects.locations.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters []types.Dataproc_WorkflowTemplateParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Required. WorkflowTemplate scheduling information.
	Placement types.Dataproc_WorkflowTemplatePlacement `json:"placement,omitempty" yaml:"placement,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
	Version int `json:"version,omitempty" yaml:"version,omitempty"`
}
