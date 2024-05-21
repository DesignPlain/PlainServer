package types

type Dataflow_PipelineWorkloadDataflowFlexTemplateRequest struct {
	// The ID of the Cloud Platform project that the job belongs to.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// If true, the request is validated but not actually executed. Defaults to false.
	ValidateOnly bool `json:"validateOnly,omitempty" yaml:"validateOnly,omitempty"`

	/*
	   Parameter to launch a job from a Flex Template.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchflextemplateparameter
	   Structure is documented below.
	*/
	LaunchParameter Dataflow_PipelineWorkloadDataflowFlexTemplateRequestLaunchParameter `json:"launchParameter,omitempty" yaml:"launchParameter,omitempty"`

	// The regional endpoint to which to direct the request. For example, us-central1, us-west1.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
