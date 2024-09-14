package types

type Bigquery_TableRangePartitioning struct {
	/*
	   The field used to determine how to create a range-based
	   partition.
	*/
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   Information required to partition based on ranges.
	   Structure is documented below.
	*/
	Range Bigquery_TableRangePartitioningRange `json:"range,omitempty" yaml:"range,omitempty"`
}
