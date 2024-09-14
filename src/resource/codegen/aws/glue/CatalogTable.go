package glue

import types "libds/aws/types"

type CatalogTable struct {
	// Properties associated with this table, as a list of key-value pairs.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor). See `storage_descriptor` below.
	StorageDescriptor types.Glue_CatalogTableStorageDescriptor `json:"storageDescriptor,omitempty" yaml:"storageDescriptor,omitempty"`

	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText string `json:"viewExpandedText,omitempty" yaml:"viewExpandedText,omitempty"`

	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText string `json:"viewOriginalText,omitempty" yaml:"viewOriginalText,omitempty"`

	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Name of the table. For Hive compatibility, this must be entirely lowercase.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Retention time for this table.
	Retention int `json:"retention,omitempty" yaml:"retention,omitempty"`

	// Owner of the table.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`

	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partition_keys` below.
	PartitionKeys []types.Glue_CatalogTablePartitionKey `json:"partitionKeys,omitempty" yaml:"partitionKeys,omitempty"`

	/*
	   Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.

	   The follow arguments are optional:
	*/
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Description of the table.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Configuration block for open table formats. See `open_table_format_input` below.
	OpenTableFormatInput types.Glue_CatalogTableOpenTableFormatInput `json:"openTableFormatInput,omitempty" yaml:"openTableFormatInput,omitempty"`

	// Configuration block for a maximum of 3 partition indexes. See `partition_index` below.
	PartitionIndices []types.Glue_CatalogTablePartitionIndex `json:"partitionIndices,omitempty" yaml:"partitionIndices,omitempty"`

	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
	TableType string `json:"tableType,omitempty" yaml:"tableType,omitempty"`

	// Configuration block of a target table for resource linking. See `target_table` below.
	TargetTable types.Glue_CatalogTableTargetTable `json:"targetTable,omitempty" yaml:"targetTable,omitempty"`
}
