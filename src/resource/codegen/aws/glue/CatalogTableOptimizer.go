package glue

import types "libds/aws/types"

type CatalogTableOptimizer struct {
	// The name of the database in the catalog in which the table resides.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The name of the table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// The type of table optimizer. Currently, the only valid value is compaction.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The Catalog ID of the table.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// A configuration block that defines the table optimizer settings. The block contains:
	Configuration types.Glue_CatalogTableOptimizerConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`
}
