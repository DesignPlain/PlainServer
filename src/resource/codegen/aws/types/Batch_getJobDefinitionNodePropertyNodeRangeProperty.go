package types

type Batch_getJobDefinitionNodePropertyNodeRangeProperty struct {
	// The container details for the node range.
	Containers []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainer `json:"containers,omitempty" yaml:"containers,omitempty"`

	// The range of nodes, using node index values. A range of 0:3 indicates nodes with index values of 0 through 3. I
	TargetNodes string `json:"targetNodes,omitempty" yaml:"targetNodes,omitempty"`
}
