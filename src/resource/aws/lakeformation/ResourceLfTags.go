package lakeformation

import types "DesignSphere_Server/src/resource/aws/types"

type ResourceLfTags struct {
	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Configuration block for a database resource. See below.
	Database types.Lakeformation_ResourceLfTagsDatabase `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   Set of LF-tags to attach to the resource. See below.

	   Exactly one of the following is required:
	*/
	LfTags []types.Lakeformation_ResourceLfTagsLfTag `json:"lfTags,omitempty" yaml:"lfTags,omitempty"`

	// Configuration block for a table resource. See below.
	Table types.Lakeformation_ResourceLfTagsTable `json:"table,omitempty" yaml:"table,omitempty"`

	/*
	   Configuration block for a table with columns resource. See below.

	   The following arguments are optional:
	*/
	TableWithColumns types.Lakeformation_ResourceLfTagsTableWithColumns `json:"tableWithColumns,omitempty" yaml:"tableWithColumns,omitempty"`
}
