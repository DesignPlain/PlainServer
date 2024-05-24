package types

type Timestreamwrite_TableSchema struct {
	// A non-empty list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed. See Composite Partition Key below for more details.
	CompositePartitionKey Timestreamwrite_TableSchemaCompositePartitionKey `json:"compositePartitionKey,omitempty" yaml:"compositePartitionKey,omitempty"`
}
