package types

type Glue_CatalogTableStorageDescriptorColumn struct {
	// Free-form text comment.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Name of the Column.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value pairs defining properties associated with the column.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Datatype of data in the Column.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
