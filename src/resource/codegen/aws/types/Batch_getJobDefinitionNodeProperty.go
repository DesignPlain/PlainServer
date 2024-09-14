package types

type Batch_getJobDefinitionNodeProperty struct {
	// Specifies the node index for the main node of a multi-node parallel job. This node index value must be fewer than the number of nodes.
	MainNode int `json:"mainNode,omitempty" yaml:"mainNode,omitempty"`

	// A list of node ranges and their properties that are associated with a multi-node parallel job.
	NodeRangeProperties []Batch_getJobDefinitionNodePropertyNodeRangeProperty `json:"nodeRangeProperties,omitempty" yaml:"nodeRangeProperties,omitempty"`

	// The number of nodes that are associated with a multi-node parallel job.
	NumNodes int `json:"numNodes,omitempty" yaml:"numNodes,omitempty"`
}
