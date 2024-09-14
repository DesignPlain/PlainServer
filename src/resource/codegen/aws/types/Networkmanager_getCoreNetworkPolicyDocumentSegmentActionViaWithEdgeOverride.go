package types

type Networkmanager_getCoreNetworkPolicyDocumentSegmentActionViaWithEdgeOverride struct {
	// A list of strings. The list of edges associated with the network function group.
	EdgeSets []string `json:"edgeSets,omitempty" yaml:"edgeSets,omitempty"`

	// The preferred edge to use.
	UseEdge string `json:"useEdge,omitempty" yaml:"useEdge,omitempty"`
}
