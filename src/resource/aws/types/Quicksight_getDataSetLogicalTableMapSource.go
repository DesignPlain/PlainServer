package types

type Quicksight_getDataSetLogicalTableMapSource struct {
	//
	PhysicalTableId string `json:"physicalTableId,omitempty" yaml:"physicalTableId,omitempty"`

	//
	DataSetArn string `json:"dataSetArn,omitempty" yaml:"dataSetArn,omitempty"`

	//
	JoinInstructions []Quicksight_getDataSetLogicalTableMapSourceJoinInstruction `json:"joinInstructions,omitempty" yaml:"joinInstructions,omitempty"`
}
