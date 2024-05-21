package types

type Logging_getSinkBigqueryOption struct {
	// Whether [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables) are used.
	UsePartitionedTables bool `json:"usePartitionedTables,omitempty" yaml:"usePartitionedTables,omitempty"`
}
