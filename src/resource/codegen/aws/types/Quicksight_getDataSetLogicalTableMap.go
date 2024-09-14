package types

type Quicksight_getDataSetLogicalTableMap struct {
	//
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	//
	DataTransforms []Quicksight_getDataSetLogicalTableMapDataTransform `json:"dataTransforms,omitempty" yaml:"dataTransforms,omitempty"`

	//
	LogicalTableMapId string `json:"logicalTableMapId,omitempty" yaml:"logicalTableMapId,omitempty"`

	//
	Sources []Quicksight_getDataSetLogicalTableMapSource `json:"sources,omitempty" yaml:"sources,omitempty"`
}
