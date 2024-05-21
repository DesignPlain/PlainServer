package types

type Compute_RouterStatusBestRoutesForRouter struct {
	//
	SelfLink string `json:"selfLink,omitempty" yaml:"selfLink,omitempty"`

	/*
	   The destination range of outgoing packets that this route applies to.
	   Only IPv4 is supported.
	*/
	DestRange string `json:"destRange,omitempty" yaml:"destRange,omitempty"`

	/*
	   URL to an instance that should handle matching packets.
	   You can specify this as a full or partial URL. For example:
	   - 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance'
	   - 'projects/project/zones/zone/instances/instance'
	   - 'zones/zone/instances/instance'
	   - Just the instance name, with the zone in 'next_hop_instance_zone'.
	*/
	NextHopInstance string `json:"nextHopInstance,omitempty" yaml:"nextHopInstance,omitempty"`

	// The zone of the instance specified in next_hop_instance. Omit if next_hop_instance is specified as a URL.
	NextHopInstanceZone string `json:"nextHopInstanceZone,omitempty" yaml:"nextHopInstanceZone,omitempty"`

	// Network IP address of an instance that should handle matching packets.
	NextHopIp string `json:"nextHopIp,omitempty" yaml:"nextHopIp,omitempty"`

	// URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel string `json:"nextHopVpnTunnel,omitempty" yaml:"nextHopVpnTunnel,omitempty"`

	/*
	   URL to a gateway that should handle matching packets.
	   Currently, you can only specify the internet gateway, using a full or
	   partial valid URL:
	   - 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
	   - 'projects/project/global/gateways/default-internet-gateway'
	   - 'global/gateways/default-internet-gateway'
	   - The string 'default-internet-gateway'.
	*/
	NextHopGateway string `json:"nextHopGateway,omitempty" yaml:"nextHopGateway,omitempty"`

	/*
	   The IP address or URL to a forwarding rule of type
	   loadBalancingScheme=INTERNAL that should handle matching
	   packets.

	   With the GA provider you can only specify the forwarding
	   rule as a partial or full URL. For example, the following
	   are all valid values:
	   - 10.128.0.56
	   - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	   - regions/region/forwardingRules/forwardingRule

	   When the beta provider, you can also specify the IP address
	   of a forwarding rule from the same VPC or any peered VPC.

	   Note that this can only be used when the destinationRange is
	   a public (non-RFC 1918) IP CIDR range.
	*/
	NextHopIlb string `json:"nextHopIlb,omitempty" yaml:"nextHopIlb,omitempty"`

	/*
	   An optional description of this resource. Provide this property
	   when you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The network name or resource link to the parent
	   network of this subnetwork.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// URL to a Network that should handle matching packets.
	NextHopNetwork string `json:"nextHopNetwork,omitempty" yaml:"nextHopNetwork,omitempty"`

	/*
	   The priority of this route. Priority is used to break ties in cases
	   where there is more than one matching route of equal prefix length.

	   In the case of two routes with equal prefix length, the one with the
	   lowest-numbered priority value wins.

	   Default value is 1000. Valid range is 0 through 65535.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   The ID of the project in which the resource
	   belongs. If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A list of instance tags to which this route applies.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the router.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
