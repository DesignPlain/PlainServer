package types

type Quicksight_DataSetLogicalTableMapSource struct {
	// ARN of the parent data set.
	DataSetArn string `json:"dataSetArn,omitempty" yaml:"dataSetArn,omitempty"`

	// Specifies the result of a join of two logical tables. See join_instruction.
	JoinInstruction Quicksight_DataSetLogicalTableMapSourceJoinInstruction `json:"joinInstruction,omitempty" yaml:"joinInstruction,omitempty"`

	// Physical table ID.
	PhysicalTableId string `json:"physicalTableId,omitempty" yaml:"physicalTableId,omitempty"`
}
