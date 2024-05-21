package types



type Dataflow_PipelineWorkloadDataflowLaunchTemplateRequest struct {
	// The ID of the Cloud Platform project that the job belongs to.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// (Optional)
	ValidateOnly bool `json:"validateOnly,omitempty" yaml:"validateOnly,omitempty"`

	// A Cloud Storage path to the template from which to create the job. Must be a valid Cloud Storage URL, beginning with 'gs://'.
	GcsPath string `json:"gcsPath,omitempty" yaml:"gcsPath,omitempty"`

	/*
	   The parameters of the template to launch. This should be part of the body of the POST request.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchtemplateparameters
	   Structure is documented below.
	*/
	LaunchParameters Dataflow_PipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters `json:"launchParameters,omitempty" yaml:"launchParameters,omitempty"`

	// The regional endpoint to which to direct the request.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
