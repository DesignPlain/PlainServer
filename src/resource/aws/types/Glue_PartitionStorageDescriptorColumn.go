package types

type Glue_PartitionStorageDescriptorColumn struct {
	// Free-form text comment.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The name of the Column.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The datatype of data in the Column.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
