package types

type Glue_getCatalogTableStorageDescriptorColumn struct {
	// Datatype of data in the Column.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Free-form text comment.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Name of the table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
