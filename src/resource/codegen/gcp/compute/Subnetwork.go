package compute

import types "libds/gcp/types"

type Subnetwork struct {
	/*
	   An array of configurations for secondary IP ranges for VM instances
	   contained in this subnetwork. The primary IP of such VM must belong
	   to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	   to either primary or secondary ranges.
	   Structure is documented below.
	*/
	SecondaryIpRanges []types.Compute_SubnetworkSecondaryIpRange `json:"secondaryIpRanges,omitempty" yaml:"secondaryIpRanges,omitempty"`

	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIpv6Prefix string `json:"externalIpv6Prefix,omitempty" yaml:"externalIpv6Prefix,omitempty"`

	/*
	   The range of internal addresses that are owned by this subnetwork.
	   Provide this property when you create the subnetwork. For example,
	   10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	   non-overlapping within a network. Only IPv4 is supported.
	*/
	IpCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`

	/*
	   The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation
	   or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	   cannot enable direct path.
	   Possible values are: `EXTERNAL`, `INTERNAL`.
	*/
	Ipv6AccessType string `json:"ipv6AccessType,omitempty" yaml:"ipv6AccessType,omitempty"`

	/*
	   The purpose of the resource. This field can be either `PRIVATE_RFC_1918`, `REGIONAL_MANAGED_PROXY`, `GLOBAL_MANAGED_PROXY`, `PRIVATE_SERVICE_CONNECT` or `PRIVATE_NAT`.
	   A subnet with purpose set to `REGIONAL_MANAGED_PROXY` is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
	   A subnetwork in a given region with purpose set to `GLOBAL_MANAGED_PROXY` is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
	   A subnetwork with purpose set to `PRIVATE_SERVICE_CONNECT` reserves the subnet for hosting a Private Service Connect published service.
	   A subnetwork with purpose set to `PRIVATE_NAT` is used as source range for Private NAT gateways.
	   Note that `REGIONAL_MANAGED_PROXY` is the preferred setting for all regional Envoy load balancers.
	   If unspecified, the purpose defaults to `PRIVATE_RFC_1918`.
	*/
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`

	/*
	   The role of subnetwork.
	   Currently, this field is only used when `purpose` is `REGIONAL_MANAGED_PROXY`.
	   The value can be set to `ACTIVE` or `BACKUP`.
	   An `ACTIVE` subnetwork is one that is currently being used for Envoy-based load balancers in a region.
	   A `BACKUP` subnetwork is one that is ready to be promoted to `ACTIVE` or is currently draining.
	   Possible values are: `ACTIVE`, `BACKUP`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	/*
	   The name of the resource, provided by the client when initially
	   creating the resource. The name must be 1-63 characters long, and
	   comply with RFC1035. Specifically, the name must be 1-63 characters
	   long and match the regular expression `a-z?` which
	   means the first character must be a lowercase letter, and all
	   following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   When enabled, VMs in this subnetwork without external IP addresses can
	   access Google APIs and services by using Private Google Access.
	*/
	PrivateIpGoogleAccess bool `json:"privateIpGoogleAccess,omitempty" yaml:"privateIpGoogleAccess,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The network this subnet belongs to.
	   Only networks that are in the distributed mode can have subnetworks.


	   - - -
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess string `json:"privateIpv6GoogleAccess,omitempty" yaml:"privateIpv6GoogleAccess,omitempty"`

	// The GCP region for this subnetwork.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	   If not specified IPV4_ONLY will be used.
	   Possible values are: `IPV4_ONLY`, `IPV4_IPV6`.
	*/
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`

	/*
	   Typically packets destined to IPs within the subnetwork range that do not match
	   existing resources are dropped and prevented from leaving the VPC.
	   Setting this field to true will allow these packets to match dynamic routes injected
	   via BGP even if their destinations match existing subnet ranges.
	*/
	AllowSubnetCidrRoutesOverlap bool `json:"allowSubnetCidrRoutesOverlap,omitempty" yaml:"allowSubnetCidrRoutesOverlap,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource. This field can be set only at resource
	   creation time.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   This field denotes the VPC flow logging options for this subnetwork. If
	   logging is enabled, logs are exported to Cloud Logging. Flow logging
	   isn't supported if the subnet `purpose` field is set to subnetwork is
	   `REGIONAL_MANAGED_PROXY` or `GLOBAL_MANAGED_PROXY`.
	   Structure is documented below.
	*/
	LogConfig types.Compute_SubnetworkLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`
}
