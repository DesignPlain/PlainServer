package types

type Glue_getScriptDagEdge struct {
	// ID of the node at which the edge starts.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// ID of the node at which the edge ends.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// Target of the edge.
	TargetParameter string `json:"targetParameter,omitempty" yaml:"targetParameter,omitempty"`
}
