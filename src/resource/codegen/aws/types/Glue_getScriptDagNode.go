package types

type Glue_getScriptDagNode struct {
	// Nested configuration an argument or property of a node. Defined below.
	Args []Glue_getScriptDagNodeArg `json:"args,omitempty" yaml:"args,omitempty"`

	// Node identifier that is unique within the node's graph.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Line number of the node.
	LineNumber int `json:"lineNumber,omitempty" yaml:"lineNumber,omitempty"`

	// Type of node this is.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`
}
