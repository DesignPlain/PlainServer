package types

type Quicksight_DataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties struct {
	// A value that indicates that a row in a table is uniquely identified by the columns in a join key. This is used by Amazon QuickSight to optimize query performance.
	UniqueKey bool `json:"uniqueKey,omitempty" yaml:"uniqueKey,omitempty"`
}
