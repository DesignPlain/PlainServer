package types

type Vmwareengine_getPrivateCloudNetworkConfig struct {
	// DNS Server IP of the Private Cloud.
	DnsServerIp string `json:"dnsServerIp,omitempty" yaml:"dnsServerIp,omitempty"`

	// Management CIDR used by VMware management appliances.
	ManagementCidr string `json:"managementCidr,omitempty" yaml:"managementCidr,omitempty"`

	/*
	   The IP address layout version of the management IP address range.
	   Possible versions include:
	   - managementIpAddressLayoutVersion=1: Indicates the legacy IP address layout used by some existing private clouds. This is no longer supported for new private clouds
	   as it does not support all features.
	   - managementIpAddressLayoutVersion=2: Indicates the latest IP address layout
	   used by all newly created private clouds. This version supports all current features.
	*/
	ManagementIpAddressLayoutVersion int `json:"managementIpAddressLayoutVersion,omitempty" yaml:"managementIpAddressLayoutVersion,omitempty"`

	/*
	   The relative resource name of the VMware Engine network attached to the private cloud.
	   Specify the name in the following form: projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
	   where {project} can either be a project number or a project ID.
	*/
	VmwareEngineNetwork string `json:"vmwareEngineNetwork,omitempty" yaml:"vmwareEngineNetwork,omitempty"`

	/*
	   The canonical name of the VMware Engine network in
	   the form: projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
	*/
	VmwareEngineNetworkCanonical string `json:"vmwareEngineNetworkCanonical,omitempty" yaml:"vmwareEngineNetworkCanonical,omitempty"`
}
