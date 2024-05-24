package types

type Keyspaces_TableSchemaDefinition struct {
	// The columns that are part of the partition key of the table .
	PartitionKeys []Keyspaces_TableSchemaDefinitionPartitionKey `json:"partitionKeys,omitempty" yaml:"partitionKeys,omitempty"`

	// The columns that have been defined as `STATIC`. Static columns store values that are shared by all rows in the same partition.
	StaticColumns []Keyspaces_TableSchemaDefinitionStaticColumn `json:"staticColumns,omitempty" yaml:"staticColumns,omitempty"`

	// The columns that are part of the clustering key of the table.
	ClusteringKeys []Keyspaces_TableSchemaDefinitionClusteringKey `json:"clusteringKeys,omitempty" yaml:"clusteringKeys,omitempty"`

	// The regular columns of the table.
	Columns []Keyspaces_TableSchemaDefinitionColumn `json:"columns,omitempty" yaml:"columns,omitempty"`
}
