package vmwareengine

type NetworkPeering struct {
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutes bool `json:"importCustomRoutes,omitempty" yaml:"importCustomRoutes,omitempty"`

	/*
	   The ID of the Network Peering.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The type of the network to peer with the VMware Engine network.
	   Possible values are: `STANDARD`, `VMWARE_ENGINE_NETWORK`, `PRIVATE_SERVICES_ACCESS`, `NETAPP_CLOUD_VOLUMES`, `THIRD_PARTY_SERVICE`, `DELL_POWERSCALE`.
	*/
	PeerNetworkType string `json:"peerNetworkType,omitempty" yaml:"peerNetworkType,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// User-provided description for this network peering.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// True if all subnet routes with a public IP address range are exported; false otherwise.
	ExportCustomRoutesWithPublicIp bool `json:"exportCustomRoutesWithPublicIp,omitempty" yaml:"exportCustomRoutesWithPublicIp,omitempty"`

	/*
	   The relative resource name of the network to peer with a standard VMware Engine network.
	   The provided network can be a consumer VPC network or another standard VMware Engine network.
	*/
	PeerNetwork string `json:"peerNetwork,omitempty" yaml:"peerNetwork,omitempty"`

	/*
	   The relative resource name of the VMware Engine network. Specify the name in the following form:
	   projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}
	   can either be a project number or a project ID.
	*/
	VmwareEngineNetwork string `json:"vmwareEngineNetwork,omitempty" yaml:"vmwareEngineNetwork,omitempty"`

	// True if custom routes are exported to the peered network; false otherwise.
	ExportCustomRoutes bool `json:"exportCustomRoutes,omitempty" yaml:"exportCustomRoutes,omitempty"`

	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutesWithPublicIp bool `json:"importCustomRoutesWithPublicIp,omitempty" yaml:"importCustomRoutesWithPublicIp,omitempty"`
}
