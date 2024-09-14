package types

type Lakeformation_PermissionsLfTag struct {
	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// The key-name for the tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   A list of possible values an attribute can take.

	   The following argument is optional:
	*/
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
