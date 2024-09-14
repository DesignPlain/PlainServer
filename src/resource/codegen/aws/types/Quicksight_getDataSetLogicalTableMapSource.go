package types

type Quicksight_getDataSetLogicalTableMapSource struct {
	//
	DataSetArn string `json:"dataSetArn,omitempty" yaml:"dataSetArn,omitempty"`

	//
	JoinInstructions []Quicksight_getDataSetLogicalTableMapSourceJoinInstruction `json:"joinInstructions,omitempty" yaml:"joinInstructions,omitempty"`

	//
	PhysicalTableId string `json:"physicalTableId,omitempty" yaml:"physicalTableId,omitempty"`
}
