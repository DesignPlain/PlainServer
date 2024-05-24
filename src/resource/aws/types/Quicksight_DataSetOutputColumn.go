package types

type Quicksight_DataSetOutputColumn struct {
	// Field folder description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Display name for the dataset.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Data type of the column.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
