package types

type Dataflow_PipelineWorkloadDataflowFlexTemplateRequestLaunchParameter struct {
	// Cloud Storage path to a file with a JSON-serialized ContainerSpec as content.
	ContainerSpecGcsPath string `json:"containerSpecGcsPath,omitempty" yaml:"containerSpecGcsPath,omitempty"`

	/*
	   The runtime environment for the Flex Template job.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#FlexTemplateRuntimeEnvironment
	   Structure is documented below.
	*/
	Environment Dataflow_PipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment `json:"environment,omitempty" yaml:"environment,omitempty"`

	// The job name to use for the created job. For an update job request, the job name should be the same as the existing running job.
	JobName string `json:"jobName,omitempty" yaml:"jobName,omitempty"`

	/*
	   Launch options for this Flex Template job. This is a common set of options across languages and templates. This should not be used to pass job parameters.
	   'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	*/
	LaunchOptions map[string]string `json:"launchOptions,omitempty" yaml:"launchOptions,omitempty"`

	/*
	   'The parameters for the Flex Template. Example: {"numWorkers":"5"}'
	   'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	*/
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	/*
	   'Use this to pass transform name mappings for streaming update jobs. Example: {"oldTransformName":"newTransformName",...}'
	   'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	*/
	TransformNameMappings map[string]string `json:"transformNameMappings,omitempty" yaml:"transformNameMappings,omitempty"`

	// Set this to true if you are sending a request to update a running streaming job. When set, the job name should be the same as the running job.
	Update bool `json:"update,omitempty" yaml:"update,omitempty"`
}
