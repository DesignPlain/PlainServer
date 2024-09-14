package types

type Lakeformation_getPermissionsTable struct {
	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	/*
	   Name of the database for the table. Unique to a Data Catalog.

	   The following arguments are optional:
	*/
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Name of the table. At least one of `name` or `wildcard` is required.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether to use a wildcard representing every table under a database. At least one of `name` or `wildcard` is required. Defaults to `false`.
	Wildcard bool `json:"wildcard,omitempty" yaml:"wildcard,omitempty"`
}
