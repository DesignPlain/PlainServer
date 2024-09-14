package types

type Dataproc_WorkflowTemplateJobSparkSqlJob struct {
	// The HCFS URI of the script that contains SQL queries.
	QueryFileUri string `json:"queryFileUri,omitempty" yaml:"queryFileUri,omitempty"`

	// A list of queries.
	QueryList Dataproc_WorkflowTemplateJobSparkSqlJobQueryList `json:"queryList,omitempty" yaml:"queryList,omitempty"`

	// Mapping of query variable names to values (equivalent to the Spark SQL command: SET `name="value";`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" yaml:"scriptVariables,omitempty"`

	// HCFS URIs of jar files to be added to the Spark CLASSPATH.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime log config for job execution.
	LoggingConfig Dataproc_WorkflowTemplateJobSparkSqlJobLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Dataproc API may be overwritten.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
