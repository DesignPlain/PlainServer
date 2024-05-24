package types

type Emrcontainers_JobTemplateJobTemplateDataJobDriverSparkSqlJobDriver struct {
	// The SQL file to be executed.
	EntryPoint string `json:"entryPoint,omitempty" yaml:"entryPoint,omitempty"`

	// The Spark parameters to be included in the Spark SQL command.
	SparkSqlParameters string `json:"sparkSqlParameters,omitempty" yaml:"sparkSqlParameters,omitempty"`
}
