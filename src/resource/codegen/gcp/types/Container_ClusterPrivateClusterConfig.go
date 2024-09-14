package types

type Container_ClusterPrivateClusterConfig struct {
	// The name of the peering between this cluster and the Google owned VPC.
	PeeringName string `json:"peeringName,omitempty" yaml:"peeringName,omitempty"`

	// The internal IP address of this cluster's master endpoint.
	PrivateEndpoint string `json:"privateEndpoint,omitempty" yaml:"privateEndpoint,omitempty"`

	// Subnetwork in cluster's network where master's endpoint will be provisioned.
	PrivateEndpointSubnetwork string `json:"privateEndpointSubnetwork,omitempty" yaml:"privateEndpointSubnetwork,omitempty"`

	/*
	   The external IP address of this cluster's master endpoint.

	   !> The Google provider is unable to validate certain configurations of
	   `private_cluster_config` when `enable_private_nodes` is `false`. It's
	   recommended that you omit the block entirely if the field is not set to `true`.
	*/
	PublicEndpoint string `json:"publicEndpoint,omitempty" yaml:"publicEndpoint,omitempty"`

	/*
	   When `true`, the cluster's private
	   endpoint is used as the cluster endpoint and access through the public endpoint
	   is disabled. When `false`, either endpoint can be used. This field only applies
	   to private clusters, when `enable_private_nodes` is `true`.
	*/
	EnablePrivateEndpoint bool `json:"enablePrivateEndpoint,omitempty" yaml:"enablePrivateEndpoint,omitempty"`

	/*
	   Enables the private cluster feature,
	   creating a private endpoint on the cluster. In a private cluster, nodes only
	   have RFC 1918 private addresses and communicate with the master's private
	   endpoint via private networking.
	*/
	EnablePrivateNodes bool `json:"enablePrivateNodes,omitempty" yaml:"enablePrivateNodes,omitempty"`

	/*
	   Controls cluster master global
	   access settings. If unset, the provider will no longer manage this field and will
	   not modify the previously-set value. Structure is documented below.
	*/
	MasterGlobalAccessConfig Container_ClusterPrivateClusterConfigMasterGlobalAccessConfig `json:"masterGlobalAccessConfig,omitempty" yaml:"masterGlobalAccessConfig,omitempty"`

	/*
	   The IP range in CIDR notation to use for
	   the hosted master network. This range will be used for assigning private IP
	   addresses to the cluster master(s) and the ILB VIP. This range must not overlap
	   with any other ranges in use within the cluster's network, and it must be a /28
	   subnet. See [Private Cluster Limitations](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters#req_res_lim)
	   for more details. This field only applies to private clusters, when
	   `enable_private_nodes` is `true`.
	*/
	MasterIpv4CidrBlock string `json:"masterIpv4CidrBlock,omitempty" yaml:"masterIpv4CidrBlock,omitempty"`
}
