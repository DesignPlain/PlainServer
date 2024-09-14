package types

type Dataflow_PipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters struct {
	/*
	   Map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job. Only applicable when updating a pipeline.
	   'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	*/
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" yaml:"transformNameMapping,omitempty"`

	// If set, replace the existing pipeline with the name specified by jobName with this pipeline, preserving state.
	Update bool `json:"update,omitempty" yaml:"update,omitempty"`

	/*
	   The runtime environment for the job.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#RuntimeEnvironment
	   Structure is documented below.
	*/
	Environment Dataflow_PipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironment `json:"environment,omitempty" yaml:"environment,omitempty"`

	// The job name to use for the created job.
	JobName string `json:"jobName,omitempty" yaml:"jobName,omitempty"`

	/*
	   The runtime parameters to pass to the job.
	   'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	*/
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
