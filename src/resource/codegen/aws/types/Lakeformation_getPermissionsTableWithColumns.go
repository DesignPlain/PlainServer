package types

type Lakeformation_getPermissionsTableWithColumns struct {
	// Name of the database for the table with columns resource. Unique to the Data Catalog.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Set of column names for the table to exclude. At least one of `column_names` or `excluded_column_names` is required.
	ExcludedColumnNames []string `json:"excludedColumnNames,omitempty" yaml:"excludedColumnNames,omitempty"`

	/*
	   Name of the table resource.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether to use a wildcard representing every table under a database. At least one of `name` or `wildcard` is required. Defaults to `false`.
	Wildcard bool `json:"wildcard,omitempty" yaml:"wildcard,omitempty"`

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Set of column names for the table. At least one of `column_names` or `excluded_column_names` is required.
	ColumnNames []string `json:"columnNames,omitempty" yaml:"columnNames,omitempty"`
}
