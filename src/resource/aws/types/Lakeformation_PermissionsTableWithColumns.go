package types

type Lakeformation_PermissionsTableWithColumns struct {
	/*
	   Whether to use a column wildcard. If `excluded_column_names` is included, `wildcard` must be set to `true` to avoid the provider reporting a difference.

	   The following arguments are optional:
	*/
	Wildcard bool `json:"wildcard,omitempty" yaml:"wildcard,omitempty"`

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Set of column names for the table.
	ColumnNames []string `json:"columnNames,omitempty" yaml:"columnNames,omitempty"`

	// Name of the database for the table with columns resource. Unique to the Data Catalog.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Set of column names for the table to exclude. If `excluded_column_names` is included, `wildcard` must be set to `true` to avoid the provider reporting a difference.
	ExcludedColumnNames []string `json:"excludedColumnNames,omitempty" yaml:"excludedColumnNames,omitempty"`

	// Name of the table resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
