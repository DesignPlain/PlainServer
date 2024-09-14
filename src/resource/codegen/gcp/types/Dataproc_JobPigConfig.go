package types

type Dataproc_JobPigConfig struct {
	// Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" yaml:"continueOnFailure,omitempty"`

	/*
	   HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.

	   - `logging_config.driver_log_levels`- (Required) The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	*/
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime logging config of the job
	LoggingConfig Dataproc_JobPigConfigLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in `/etc/hadoop/conf/--site.xml`, `/etc/pig/conf/pig.properties`, and classes in user code.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	/*
	   HCFS URI of file containing Hive script to execute as the job.
	   Conflicts with `query_list`
	*/
	QueryFileUri string `json:"queryFileUri,omitempty" yaml:"queryFileUri,omitempty"`

	/*
	   The list of Hive queries or statements to execute as part of the job.
	   Conflicts with `query_file_uri`
	*/
	QueryLists []string `json:"queryLists,omitempty" yaml:"queryLists,omitempty"`

	// Mapping of query variable names to values (equivalent to the Pig command: `name=[value]`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" yaml:"scriptVariables,omitempty"`
}
