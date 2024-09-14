package types

type Quicksight_DataSetLogicalTableMapDataTransformTagColumnOperation struct {
	// Column name.
	ColumnName string `json:"columnName,omitempty" yaml:"columnName,omitempty"`

	// The dataset column tag, currently only used for geospatial type tagging. See tags.
	Tags []Quicksight_DataSetLogicalTableMapDataTransformTagColumnOperationTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
