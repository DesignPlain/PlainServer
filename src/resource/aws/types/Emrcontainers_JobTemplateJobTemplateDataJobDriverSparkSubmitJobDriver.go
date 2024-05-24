package types

type Emrcontainers_JobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver struct {
	// The entry point of job application.
	EntryPoint string `json:"entryPoint,omitempty" yaml:"entryPoint,omitempty"`

	// The arguments for job application.
	EntryPointArguments []string `json:"entryPointArguments,omitempty" yaml:"entryPointArguments,omitempty"`

	// The Spark submit parameters that are used for job runs.
	SparkSubmitParameters string `json:"sparkSubmitParameters,omitempty" yaml:"sparkSubmitParameters,omitempty"`
}
