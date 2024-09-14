package types

type Networkmanager_getCoreNetworkPolicyDocumentSegmentActionVia struct {
	// A list of strings. The network function group to use for the service insertion action.
	NetworkFunctionGroups []string `json:"networkFunctionGroups,omitempty" yaml:"networkFunctionGroups,omitempty"`

	// Any edge overrides and the preferred edge to use.
	WithEdgeOverrides []Networkmanager_getCoreNetworkPolicyDocumentSegmentActionViaWithEdgeOverride `json:"withEdgeOverrides,omitempty" yaml:"withEdgeOverrides,omitempty"`
}
