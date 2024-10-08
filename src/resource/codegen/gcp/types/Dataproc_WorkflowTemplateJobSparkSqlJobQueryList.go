package types

type Dataproc_WorkflowTemplateJobSparkSqlJobQueryList struct {
	// Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: "hiveJob": { "queryList": { "queries": } }
	Queries []string `json:"queries,omitempty" yaml:"queries,omitempty"`
}
