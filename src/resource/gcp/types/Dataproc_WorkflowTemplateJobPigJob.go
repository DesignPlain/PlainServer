package types

type Dataproc_WorkflowTemplateJobPigJob struct {
	// Mapping of query variable names to values (equivalent to the Pig command: `name=`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" yaml:"scriptVariables,omitempty"`

	// Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" yaml:"continueOnFailure,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime log config for job execution.
	LoggingConfig Dataproc_WorkflowTemplateJobPigJobLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/--site.xml, /etc/pig/conf/pig.properties, and classes in user code.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// The HCFS URI of the script that contains the Pig queries.
	QueryFileUri string `json:"queryFileUri,omitempty" yaml:"queryFileUri,omitempty"`

	// A list of queries.
	QueryList Dataproc_WorkflowTemplateJobPigJobQueryList `json:"queryList,omitempty" yaml:"queryList,omitempty"`
}
