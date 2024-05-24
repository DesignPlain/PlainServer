package lakeformation

type LfTag struct {
	// List of possible values an attribute can take.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Key-name for the tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
