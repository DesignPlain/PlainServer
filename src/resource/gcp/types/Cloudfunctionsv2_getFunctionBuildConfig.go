package types

type Cloudfunctionsv2_getFunctionBuildConfig struct {
	/*
	   The name of the function (as defined in source code) that will be executed.
	   Defaults to the resource name suffix, if not specified. For backward
	   compatibility, if function with given name is not found, then the system
	   will try to use function named "function". For Node.js this is name of a
	   function exported by the module specified in source_location.
	*/
	EntryPoint string `json:"entryPoint,omitempty" yaml:"entryPoint,omitempty"`

	// User-provided build-time environment variables for the function.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	/*
	   The runtime in which to run the function. Required when deploying a new
	   function, optional when updating an existing function.
	*/
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	// The location of the function source code.
	Sources []Cloudfunctionsv2_getFunctionBuildConfigSource `json:"sources,omitempty" yaml:"sources,omitempty"`

	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	WorkerPool string `json:"workerPool,omitempty" yaml:"workerPool,omitempty"`

	/*
	   The Cloud Build name of the latest successful
	   deployment of the function.
	*/
	Build string `json:"build,omitempty" yaml:"build,omitempty"`

	// User managed repository created in Artifact Registry optionally with a customer managed encryption key.
	DockerRepository string `json:"dockerRepository,omitempty" yaml:"dockerRepository,omitempty"`
}
