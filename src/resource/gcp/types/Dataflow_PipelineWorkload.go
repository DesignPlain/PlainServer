package types

type Dataflow_PipelineWorkload struct {
	/*
	   Template information and additional parameters needed to launch a Dataflow job using the flex launch API.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchflextemplaterequest
	   Structure is documented below.
	*/
	DataflowFlexTemplateRequest Dataflow_PipelineWorkloadDataflowFlexTemplateRequest `json:"dataflowFlexTemplateRequest,omitempty" yaml:"dataflowFlexTemplateRequest,omitempty"`

	/*
	   Template information and additional parameters needed to launch a Dataflow job using the standard launch API.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchtemplaterequest
	   Structure is documented below.
	*/
	DataflowLaunchTemplateRequest Dataflow_PipelineWorkloadDataflowLaunchTemplateRequest `json:"dataflowLaunchTemplateRequest,omitempty" yaml:"dataflowLaunchTemplateRequest,omitempty"`
}
