package vmwareengine

import types "DesignSphere_Server/src/resource/gcp/types"

type NetworkPolicy struct {
	// User-provided description for this network policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   IP address range in CIDR notation used to create internet access and external IP access.
	   An RFC 1918 CIDR block, with a "/26" prefix, is required. The range cannot overlap with any
	   prefixes either in the consumer VPC network or in use by the private clouds attached to that VPC network.
	*/
	EdgeServicesCidr string `json:"edgeServicesCidr,omitempty" yaml:"edgeServicesCidr,omitempty"`

	/*
	   Network service that allows External IP addresses to be assigned to VMware workloads.
	   This service can only be enabled when internetAccess is also enabled.
	   Structure is documented below.
	*/
	ExternalIp types.Vmwareengine_NetworkPolicyExternalIp `json:"externalIp,omitempty" yaml:"externalIp,omitempty"`

	/*
	   Network service that allows VMware workloads to access the internet.
	   Structure is documented below.
	*/
	InternetAccess types.Vmwareengine_NetworkPolicyInternetAccess `json:"internetAccess,omitempty" yaml:"internetAccess,omitempty"`

	/*
	   The resource name of the location (region) to create the new network policy in.
	   Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	   For example: projects/my-project/locations/us-central1
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the Network Policy.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The relative resource name of the VMware Engine network. Specify the name in the following form:
	   projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}
	   can either be a project number or a project ID.
	*/
	VmwareEngineNetwork string `json:"vmwareEngineNetwork,omitempty" yaml:"vmwareEngineNetwork,omitempty"`
}
