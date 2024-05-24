package types

type Quicksight_getDataSetLogicalTableMapSourceJoinInstruction struct {
	//
	RightOperand string `json:"rightOperand,omitempty" yaml:"rightOperand,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	LeftJoinKeyProperties []Quicksight_getDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperty `json:"leftJoinKeyProperties,omitempty" yaml:"leftJoinKeyProperties,omitempty"`

	//
	LeftOperand string `json:"leftOperand,omitempty" yaml:"leftOperand,omitempty"`

	//
	OnClause string `json:"onClause,omitempty" yaml:"onClause,omitempty"`

	//
	RightJoinKeyProperties []Quicksight_getDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperty `json:"rightJoinKeyProperties,omitempty" yaml:"rightJoinKeyProperties,omitempty"`
}
