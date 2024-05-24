package types

type Quicksight_getDataSetLogicalTableMapDataTransform struct {
	//
	CastColumnTypeOperations []Quicksight_getDataSetLogicalTableMapDataTransformCastColumnTypeOperation `json:"castColumnTypeOperations,omitempty" yaml:"castColumnTypeOperations,omitempty"`

	//
	CreateColumnsOperations []Quicksight_getDataSetLogicalTableMapDataTransformCreateColumnsOperation `json:"createColumnsOperations,omitempty" yaml:"createColumnsOperations,omitempty"`

	//
	FilterOperations []Quicksight_getDataSetLogicalTableMapDataTransformFilterOperation `json:"filterOperations,omitempty" yaml:"filterOperations,omitempty"`

	//
	ProjectOperations []Quicksight_getDataSetLogicalTableMapDataTransformProjectOperation `json:"projectOperations,omitempty" yaml:"projectOperations,omitempty"`

	//
	RenameColumnOperations []Quicksight_getDataSetLogicalTableMapDataTransformRenameColumnOperation `json:"renameColumnOperations,omitempty" yaml:"renameColumnOperations,omitempty"`

	//
	TagColumnOperations []Quicksight_getDataSetLogicalTableMapDataTransformTagColumnOperation `json:"tagColumnOperations,omitempty" yaml:"tagColumnOperations,omitempty"`

	//
	UntagColumnOperations []Quicksight_getDataSetLogicalTableMapDataTransformUntagColumnOperation `json:"untagColumnOperations,omitempty" yaml:"untagColumnOperations,omitempty"`
}
