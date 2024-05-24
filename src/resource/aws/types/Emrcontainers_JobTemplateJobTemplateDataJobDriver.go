package types

type Emrcontainers_JobTemplateJobTemplateDataJobDriver struct {
	// The job driver for job type.
	SparkSqlJobDriver Emrcontainers_JobTemplateJobTemplateDataJobDriverSparkSqlJobDriver `json:"sparkSqlJobDriver,omitempty" yaml:"sparkSqlJobDriver,omitempty"`

	// The job driver parameters specified for spark submit.
	SparkSubmitJobDriver Emrcontainers_JobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver `json:"sparkSubmitJobDriver,omitempty" yaml:"sparkSubmitJobDriver,omitempty"`
}
