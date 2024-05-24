package types

type Emrcontainers_JobTemplateJobTemplateData struct {
	// The tags assigned to jobs started using the job template.
	JobTags map[string]string `json:"jobTags,omitempty" yaml:"jobTags,omitempty"`

	// The release version of Amazon EMR.
	ReleaseLabel string `json:"releaseLabel,omitempty" yaml:"releaseLabel,omitempty"`

	// The configuration settings that are used to override defaults configuration.
	ConfigurationOverrides Emrcontainers_JobTemplateJobTemplateDataConfigurationOverrides `json:"configurationOverrides,omitempty" yaml:"configurationOverrides,omitempty"`

	// The execution role ARN of the job run.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// Specify the driver that the job runs on. Exactly one of the two available job drivers is required, either sparkSqlJobDriver or sparkSubmitJobDriver.
	JobDriver Emrcontainers_JobTemplateJobTemplateDataJobDriver `json:"jobDriver,omitempty" yaml:"jobDriver,omitempty"`
}
