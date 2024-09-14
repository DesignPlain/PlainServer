package types

type Bigquery_TableRangePartitioningRange struct {
	// The width of each range within the partition.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	// Start of the range partitioning, inclusive.
	Start int `json:"start,omitempty" yaml:"start,omitempty"`

	// End of the range partitioning, exclusive.
	End int `json:"end,omitempty" yaml:"end,omitempty"`
}
