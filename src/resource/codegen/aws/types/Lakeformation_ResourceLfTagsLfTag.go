package types

type Lakeformation_ResourceLfTagsLfTag struct {
	// Key name for an existing LF-tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   Value from the possible values for the LF-tag.

	   The following argument is optional:
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`
}
