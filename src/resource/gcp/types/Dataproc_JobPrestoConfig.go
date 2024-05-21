package types

type Dataproc_JobPrestoConfig struct {
	// Whether to continue executing queries if a query fails. Setting to true can be useful when executing independent parallel queries. Defaults to false.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" yaml:"continueOnFailure,omitempty"`

	// The runtime logging config of the job
	LoggingConfig Dataproc_JobPrestoConfigLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	/*
	   The format in which query output will be displayed. See the Presto documentation for supported output formats.

	   - `logging_config.driver_log_levels`- (Required) The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	*/
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`

	// A mapping of property names to values. Used to set Presto session properties Equivalent to using the --session flag in the Presto CLI.
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

	// Presto client tags to attach to this query.
	ClientTags []string `json:"clientTags,omitempty" yaml:"clientTags,omitempty"`
}
