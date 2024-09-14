package types

type Quicksight_DataSetLogicalTableMap struct {
	// A display name for the logical table.
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// Transform operations that act on this logical table. For this structure to be valid, only one of the attributes can be non-null. See data_transforms.
	DataTransforms []Quicksight_DataSetLogicalTableMapDataTransform `json:"dataTransforms,omitempty" yaml:"dataTransforms,omitempty"`

	// Key of the logical table map.
	LogicalTableMapId string `json:"logicalTableMapId,omitempty" yaml:"logicalTableMapId,omitempty"`

	// Source of this logical table. See source.
	Source Quicksight_DataSetLogicalTableMapSource `json:"source,omitempty" yaml:"source,omitempty"`
}
