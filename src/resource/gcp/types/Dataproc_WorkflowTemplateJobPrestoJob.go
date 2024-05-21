package types

type Dataproc_WorkflowTemplateJobPrestoJob struct {
	// The runtime log config for job execution.
	LoggingConfig Dataproc_WorkflowTemplateJobPrestoJobLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// The format in which query output will be displayed. See the Presto documentation for supported output formats
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`

	// A mapping of property names to values. Used to set Presto (https://prestodb.io/docs/current/sql/set-session.html) Equivalent to using the --session flag in the Presto CLI
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// The HCFS URI of the script that contains SQL queries.
	QueryFileUri string `json:"queryFileUri,omitempty" yaml:"queryFileUri,omitempty"`

	// A list of queries.
	QueryList Dataproc_WorkflowTemplateJobPrestoJobQueryList `json:"queryList,omitempty" yaml:"queryList,omitempty"`

	// Presto client tags to attach to this query
	ClientTags []string `json:"clientTags,omitempty" yaml:"clientTags,omitempty"`

	// Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" yaml:"continueOnFailure,omitempty"`
}
