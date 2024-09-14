package types

type Dataproc_WorkflowTemplateJobHiveJob struct {
	// Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" yaml:"continueOnFailure,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/--site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// The HCFS URI of the script that contains Hive queries.
	QueryFileUri string `json:"queryFileUri,omitempty" yaml:"queryFileUri,omitempty"`

	// A list of queries.
	QueryList Dataproc_WorkflowTemplateJobHiveJobQueryList `json:"queryList,omitempty" yaml:"queryList,omitempty"`

	// Mapping of query variable names to values (equivalent to the Hive command: `SET name="value";`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" yaml:"scriptVariables,omitempty"`
}
