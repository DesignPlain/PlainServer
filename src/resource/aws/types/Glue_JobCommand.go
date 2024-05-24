package types

type Glue_JobCommand struct {
	// The name of the job command. Defaults to `glueetl`. Use `pythonshell` for Python Shell Job Type, `glueray` for Ray Job Type, or `gluestreaming` for Streaming Job Type. `max_capacity` needs to be set if `pythonshell` is chosen.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Python version being used to execute a Python shell job. Allowed values are 2, 3 or 3.9. Version 3 refers to Python 3.6.
	PythonVersion string `json:"pythonVersion,omitempty" yaml:"pythonVersion,omitempty"`

	// In Ray jobs, runtime is used to specify the versions of Ray, Python and additional libraries available in your environment. This field is not used in other job  For supported runtime environment values, see [Working with Ray jobs](https://docs.aws.amazon.com/glue/latest/dg/ray-jobs-section.html#author-job-ray-runtimes) in the Glue Developer Guide.
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	// Specifies the S3 path to a script that executes a job.
	ScriptLocation string `json:"scriptLocation,omitempty" yaml:"scriptLocation,omitempty"`
}
