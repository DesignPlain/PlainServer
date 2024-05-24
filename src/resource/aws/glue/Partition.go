package glue

import types "DesignSphere_Server/src/resource/aws/types"

type Partition struct {
	// The values that define the partition.
	PartitionValues []string `json:"partitionValues,omitempty" yaml:"partitionValues,omitempty"`

	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor types.Glue_PartitionStorageDescriptor `json:"storageDescriptor,omitempty" yaml:"storageDescriptor,omitempty"`

	//
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Properties associated with this table, as a list of key-value pairs.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
