package types

type Lakeformation_getPermissionsDataCellsFilter struct {
	// The name of the database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The name of the data cells filter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ID of the Data Catalog.
	TableCatalogId string `json:"tableCatalogId,omitempty" yaml:"tableCatalogId,omitempty"`

	// The name of the table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
