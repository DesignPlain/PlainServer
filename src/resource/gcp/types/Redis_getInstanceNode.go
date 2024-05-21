package types

type Redis_getInstanceNode struct {
	// Node identifying string. e.g. 'node-0', 'node-1'
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Location of the node.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
