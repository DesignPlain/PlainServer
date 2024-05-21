package types

type Redis_ClusterPscConnection struct {
	// Output only. The IP allocated on the consumer network for the PSC forwarding rule.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Output only. The URI of the consumer side forwarding rule. Example: projects/{projectNumOrId}/regions/us-east1/forwardingRules/{resourceId}.
	ForwardingRule string `json:"forwardingRule,omitempty" yaml:"forwardingRule,omitempty"`

	/*
	   Required. The consumer network where the network address of
	   the discovery endpoint will be reserved, in the form of
	   projects/{network_project_id_or_number}/global/networks/{network_id}.

	   - - -
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// Output only. The consumer projectId where the forwarding rule is created from.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// Output only. The PSC connection id of the forwarding rule connected to the service attachment.
	PscConnectionId string `json:"pscConnectionId,omitempty" yaml:"pscConnectionId,omitempty"`
}
