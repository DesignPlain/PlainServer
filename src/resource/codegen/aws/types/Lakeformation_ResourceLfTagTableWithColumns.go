package types

type Lakeformation_ResourceLfTagTableWithColumns struct {
	// Set of column names for the table.
	ColumnNames []string `json:"columnNames,omitempty" yaml:"columnNames,omitempty"`

	// Option to add column wildcard. See Column Wildcard for more details.
	ColumnWildcard Lakeformation_ResourceLfTagTableWithColumnsColumnWildcard `json:"columnWildcard,omitempty" yaml:"columnWildcard,omitempty"`

	// Name of the database for the table with columns resource. Unique to the Data Catalog.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	/*
	   Name of the table resource.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`
}
