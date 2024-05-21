package edgecontainer

import types "DesignSphere_Server/src/resource/gcp/types"

type VpnConnection struct {
	/*
	   Google Cloud Platform location.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The resource name of VPN connection
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The VPN connection Cloud Router name.
	Router string `json:"router,omitempty" yaml:"router,omitempty"`

	/*
	   Project detail of the VPC network. Required if VPC is in a different project than the cluster project.
	   Structure is documented below.
	*/
	VpcProject types.Edgecontainer_VpnConnectionVpcProject `json:"vpcProject,omitempty" yaml:"vpcProject,omitempty"`

	// The canonical Cluster name to connect to. It is in the form of projects/{project}/locations/{location}/clusters/{cluster}.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// Whether this VPN connection has HA enabled on cluster side. If enabled, when creating VPN connection we will attempt to use 2 ANG floating IPs.
	EnableHighAvailability bool `json:"enableHighAvailability,omitempty" yaml:"enableHighAvailability,omitempty"`

	/*
	   Labels associated with this resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   NAT gateway IP, or WAN IP address. If a customer has multiple NAT IPs, the customer needs to configure NAT such that only one external IP maps to the GMEC Anthos cluster.
	   This is empty if NAT is not used.
	*/
	NatGatewayIp string `json:"natGatewayIp,omitempty" yaml:"natGatewayIp,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The network ID of VPC to connect to.
	Vpc string `json:"vpc,omitempty" yaml:"vpc,omitempty"`
}
