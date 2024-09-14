package lakeformation

import types "libds/aws/types"

type ResourceLfTag struct {
	//
	Timeouts types.Lakeformation_ResourceLfTagTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Configuration block for a database resource. See Database for more details.
	Database types.Lakeformation_ResourceLfTagDatabase `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   Set of LF-tags to attach to the resource. See LF Tag for more details.

	   Exactly one of the following is required:
	*/
	LfTag types.Lakeformation_ResourceLfTagLfTag `json:"lfTag,omitempty" yaml:"lfTag,omitempty"`

	// Configuration block for a table resource. See Table for more details.
	Table types.Lakeformation_ResourceLfTagTable `json:"table,omitempty" yaml:"table,omitempty"`

	/*
	   Configuration block for a table with columns resource. See Table With Columns for more details.

	   The following arguments are optional:
	*/
	TableWithColumns types.Lakeformation_ResourceLfTagTableWithColumns `json:"tableWithColumns,omitempty" yaml:"tableWithColumns,omitempty"`
}
