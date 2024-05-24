package types

type Quicksight_DataSetLogicalTableMapDataTransform struct {
	// A transform operation that removes tags associated with a column. See untag_column_operation.
	UntagColumnOperation Quicksight_DataSetLogicalTableMapDataTransformUntagColumnOperation `json:"untagColumnOperation,omitempty" yaml:"untagColumnOperation,omitempty"`

	// A transform operation that casts a column to a different type. See cast_column_type_operation.
	CastColumnTypeOperation Quicksight_DataSetLogicalTableMapDataTransformCastColumnTypeOperation `json:"castColumnTypeOperation,omitempty" yaml:"castColumnTypeOperation,omitempty"`

	// An operation that creates calculated columns. Columns created in one such operation form a lexical closure. See create_columns_operation.
	CreateColumnsOperation Quicksight_DataSetLogicalTableMapDataTransformCreateColumnsOperation `json:"createColumnsOperation,omitempty" yaml:"createColumnsOperation,omitempty"`

	// An operation that filters rows based on some condition. See filter_operation.
	FilterOperation Quicksight_DataSetLogicalTableMapDataTransformFilterOperation `json:"filterOperation,omitempty" yaml:"filterOperation,omitempty"`

	// An operation that projects columns. Operations that come after a projection can only refer to projected columns. See project_operation.
	ProjectOperation Quicksight_DataSetLogicalTableMapDataTransformProjectOperation `json:"projectOperation,omitempty" yaml:"projectOperation,omitempty"`

	// An operation that renames a column. See rename_column_operation.
	RenameColumnOperation Quicksight_DataSetLogicalTableMapDataTransformRenameColumnOperation `json:"renameColumnOperation,omitempty" yaml:"renameColumnOperation,omitempty"`

	// An operation that tags a column with additional information. See tag_column_operation.
	TagColumnOperation Quicksight_DataSetLogicalTableMapDataTransformTagColumnOperation `json:"tagColumnOperation,omitempty" yaml:"tagColumnOperation,omitempty"`
}
