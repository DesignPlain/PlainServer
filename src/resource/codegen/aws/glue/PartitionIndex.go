package glue

import types "libds/aws/types"

type PartitionIndex struct {
	// The catalog ID where the table resides.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Configuration block for a partition index. See `partition_index` below.
	PartitionIndex types.Glue_PartitionIndexPartitionIndex `json:"partitionIndex,omitempty" yaml:"partitionIndex,omitempty"`

	// Name of the table. For Hive compatibility, this must be entirely lowercase.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
