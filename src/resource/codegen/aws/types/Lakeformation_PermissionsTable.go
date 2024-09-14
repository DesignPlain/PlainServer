package types

type Lakeformation_PermissionsTable struct {
	// Name of the table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Whether to use a wildcard representing every table under a database. Defaults to `false`.

	   The following arguments are optional:
	*/
	Wildcard bool `json:"wildcard,omitempty" yaml:"wildcard,omitempty"`

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Name of the database for the table. Unique to a Data Catalog.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`
}
