package types

type Glue_CatalogTablePartitionKey struct {
	// Free-form text comment.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Name of the Partition Key.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Datatype of data in the Partition Key.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
