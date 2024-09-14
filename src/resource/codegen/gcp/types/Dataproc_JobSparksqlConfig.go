package types

type Dataproc_JobSparksqlConfig struct {
	/*
	   HCFS URIs of jar files to be added to the Spark CLASSPATH.

	   - `logging_config.driver_log_levels`- (Required) The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	*/
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime logging config of the job
	LoggingConfig Dataproc_JobSparksqlConfigLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	/*
	   The HCFS URI of the script that contains SQL queries.
	   Conflicts with `query_list`
	*/
	QueryFileUri string `json:"queryFileUri,omitempty" yaml:"queryFileUri,omitempty"`

	/*
	   The list of SQL queries or statements to execute as part of the job.
	   Conflicts with `query_file_uri`
	*/
	QueryLists []string `json:"queryLists,omitempty" yaml:"queryLists,omitempty"`

	// Mapping of query variable names to values (equivalent to the Spark SQL command: `SET name="value";`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" yaml:"scriptVariables,omitempty"`
}
