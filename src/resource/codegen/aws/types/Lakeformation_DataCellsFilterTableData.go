package types

type Lakeformation_DataCellsFilterTableData struct {
	// A list of column names and/or nested column attributes.
	ColumnNames []string `json:"columnNames,omitempty" yaml:"columnNames,omitempty"`

	// A wildcard with exclusions. See Column Wildcard below for details.
	ColumnWildcard Lakeformation_DataCellsFilterTableDataColumnWildcard `json:"columnWildcard,omitempty" yaml:"columnWildcard,omitempty"`

	// The name of the database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The name of the data cells filter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A PartiQL predicate. See Row Filter below for details.
	RowFilter Lakeformation_DataCellsFilterTableDataRowFilter `json:"rowFilter,omitempty" yaml:"rowFilter,omitempty"`

	// The ID of the Data Catalog.
	TableCatalogId string `json:"tableCatalogId,omitempty" yaml:"tableCatalogId,omitempty"`

	// The name of the table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// ID of the data cells filter version.
	VersionId string `json:"versionId,omitempty" yaml:"versionId,omitempty"`
}
