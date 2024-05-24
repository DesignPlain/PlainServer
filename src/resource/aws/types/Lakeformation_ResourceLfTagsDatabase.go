package types

type Lakeformation_ResourceLfTagsDatabase struct {
	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	/*
	   Name of the database resource. Unique to the Data Catalog.

	   The following argument is optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
