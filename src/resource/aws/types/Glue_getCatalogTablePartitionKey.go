package types

type Glue_getCatalogTablePartitionKey struct {
	// Free-form text comment.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Name of the table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Datatype of data in the Column.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
