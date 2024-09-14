package types

type Networkmanager_getCoreNetworkPolicyDocumentNetworkFunctionGroup struct {
	// Optional description of the network function group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// This identifies the network function group container.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// This will be either `true`, that attachment acceptance is required, or `false`, that it is not required.
	RequireAttachmentAcceptance bool `json:"requireAttachmentAcceptance,omitempty" yaml:"requireAttachmentAcceptance,omitempty"`
}
