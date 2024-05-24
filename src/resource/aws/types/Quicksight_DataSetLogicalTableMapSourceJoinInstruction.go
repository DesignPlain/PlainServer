package types

type Quicksight_DataSetLogicalTableMapSourceJoinInstruction struct {
	// Join instructions provided in the ON clause of a join.
	OnClause string `json:"onClause,omitempty" yaml:"onClause,omitempty"`

	// Join key properties of the right operand. See right_join_key_properties.
	RightJoinKeyProperties Quicksight_DataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties `json:"rightJoinKeyProperties,omitempty" yaml:"rightJoinKeyProperties,omitempty"`

	// Operand on the right side of a join.
	RightOperand string `json:"rightOperand,omitempty" yaml:"rightOperand,omitempty"`

	// Type of join. Valid values are `INNER`, `OUTER`, `LEFT`, and `RIGHT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Join key properties of the left operand. See left_join_key_properties.
	LeftJoinKeyProperties Quicksight_DataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties `json:"leftJoinKeyProperties,omitempty" yaml:"leftJoinKeyProperties,omitempty"`

	// Operand on the left side of a join.
	LeftOperand string `json:"leftOperand,omitempty" yaml:"leftOperand,omitempty"`
}
