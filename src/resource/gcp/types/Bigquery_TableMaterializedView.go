package types

type Bigquery_TableMaterializedView struct {
	/*
	   Allow non incremental materialized view definition.
	   The default value is false.
	*/
	AllowNonIncrementalDefinition bool `json:"allowNonIncrementalDefinition,omitempty" yaml:"allowNonIncrementalDefinition,omitempty"`

	/*
	   Specifies whether to use BigQuery's automatic refresh for this materialized view when the base table is updated.
	   The default value is true.
	*/
	EnableRefresh bool `json:"enableRefresh,omitempty" yaml:"enableRefresh,omitempty"`

	// A query whose result is persisted.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`

	/*
	   The maximum frequency at which this materialized view will be refreshed.
	   The default value is 1800000
	*/
	RefreshIntervalMs int `json:"refreshIntervalMs,omitempty" yaml:"refreshIntervalMs,omitempty"`
}
