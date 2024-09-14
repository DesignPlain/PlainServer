package types

type Bigquery_TableView struct {
	// A query that BigQuery executes when the view is referenced.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`

	/*
	   Specifies whether to use BigQuery's legacy SQL for this view.
	   The default value is true. If set to false, the view will use BigQuery's standard SQL.
	*/
	UseLegacySql bool `json:"useLegacySql,omitempty" yaml:"useLegacySql,omitempty"`
}
