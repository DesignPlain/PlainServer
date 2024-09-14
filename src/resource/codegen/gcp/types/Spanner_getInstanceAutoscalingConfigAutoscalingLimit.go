package types

type Spanner_getInstanceAutoscalingConfigAutoscalingLimit struct {
	/*
	   Specifies number of nodes allocated to the instance. If set, this number
	   should be greater than or equal to 1.
	*/
	MinNodes int `json:"minNodes,omitempty" yaml:"minNodes,omitempty"`

	/*
	   Specifies minimum number of processing units allocated to the instance.
	   If set, this number should be multiples of 1000.
	*/
	MinProcessingUnits int `json:"minProcessingUnits,omitempty" yaml:"minProcessingUnits,omitempty"`

	/*
	   Specifies maximum number of nodes allocated to the instance. If set, this number
	   should be greater than or equal to min_nodes.
	*/
	MaxNodes int `json:"maxNodes,omitempty" yaml:"maxNodes,omitempty"`

	/*
	   Specifies maximum number of processing units allocated to the instance.
	   If set, this number should be multiples of 1000 and be greater than or equal to
	   min_processing_units.
	*/
	MaxProcessingUnits int `json:"maxProcessingUnits,omitempty" yaml:"maxProcessingUnits,omitempty"`
}
