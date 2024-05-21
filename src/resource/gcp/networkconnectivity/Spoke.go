package networkconnectivity

import types "DesignSphere_Server/src/resource/gcp/types"

type Spoke struct {
	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Immutable. The name of the spoke. Spoke names must be unique.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Immutable. The URI of the hub that this spoke is attached to.
	Hub string `json:"hub,omitempty" yaml:"hub,omitempty"`

	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments types.Networkconnectivity_SpokeLinkedInterconnectAttachments `json:"linkedInterconnectAttachments,omitempty" yaml:"linkedInterconnectAttachments,omitempty"`

	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances types.Networkconnectivity_SpokeLinkedRouterApplianceInstances `json:"linkedRouterApplianceInstances,omitempty" yaml:"linkedRouterApplianceInstances,omitempty"`

	// VPC network that is associated with the spoke.
	LinkedVpcNetwork types.Networkconnectivity_SpokeLinkedVpcNetwork `json:"linkedVpcNetwork,omitempty" yaml:"linkedVpcNetwork,omitempty"`

	// An optional description of the spoke.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels types.Networkconnectivity_SpokeLinkedVpnTunnels `json:"linkedVpnTunnels,omitempty" yaml:"linkedVpnTunnels,omitempty"`
}
