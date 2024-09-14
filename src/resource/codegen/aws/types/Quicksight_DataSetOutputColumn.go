package types

type Quicksight_DataSetOutputColumn struct {
	// Field folder description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Display name for the dataset.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
