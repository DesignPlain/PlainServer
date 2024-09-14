package compute

import types "libds/gcp/types"

type GlobalForwardingRule struct {
	/*
	   The IP protocol to which this rule applies.
	   For protocol forwarding, valid
	   options are `TCP`, `UDP`, `ESP`,
	   `AH`, `SCTP`, `ICMP` and
	   `L3_DEFAULT`.
	   The valid IP protocols are different for different load balancing products
	   as described in [Load balancing
	   features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
	   Possible values are: `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, `ICMP`.
	*/
	IpProtocol string `json:"ipProtocol,omitempty" yaml:"ipProtocol,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is created.
	   The name must be 1-63 characters long, and comply with
	   [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	   Specifically, the name must be 1-63 characters long and match the regular
	   expression `a-z?` which means the first
	   character must be a lowercase letter, and all following characters must
	   be a dash, lowercase letter, or digit, except the last character, which
	   cannot be a dash.
	   For Private Service Connect forwarding rules that forward traffic to Google
	   APIs, the forwarding rule name must be a 1-20 characters string with
	   lowercase letters and numbers and must start with a letter.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	NoAutomateDnsZone bool `json:"noAutomateDnsZone,omitempty" yaml:"noAutomateDnsZone,omitempty"`

	/*
	   The `portRange` field has the following limitations:
	   - It requires that the forwarding rule `IPProtocol` be TCP, UDP, or SCTP,
	   and
	   - It's applicable only to the following products: external passthrough
	   Network Load Balancers, internal and external proxy Network Load
	   Balancers, internal and external Application Load Balancers, external
	   protocol forwarding, and Classic VPN.
	   - Some products have restrictions on what ports can be used. See
	   [port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)
	   for details.
	   For external forwarding rules, two or more forwarding rules cannot use the
	   same `[IPAddress, IPProtocol]` pair, and cannot have overlapping
	   `portRange`s.
	   For internal forwarding rules within the same VPC network, two or more
	   forwarding rules cannot use the same `[IPAddress, IPProtocol]` pair, and
	   cannot have overlapping `portRange`s.
	*/
	PortRange string `json:"portRange,omitempty" yaml:"portRange,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	AllowPscGlobalAccess bool `json:"allowPscGlobalAccess,omitempty" yaml:"allowPscGlobalAccess,omitempty"`

	/*
	   IP address for which this forwarding rule accepts traffic. When a client
	   sends traffic to this IP address, the forwarding rule directs the traffic
	   to the referenced `target`.
	   While creating a forwarding rule, specifying an `IPAddress` is
	   required under the following circumstances:
	   - When the `target` is set to `targetGrpcProxy` and
	   `validateForProxyless` is set to `true`, the
	   `IPAddress` should be set to `0.0.0.0`.
	   - When the `target` is a Private Service Connect Google APIs
	   bundle, you must specify an `IPAddress`.

	   Otherwise, you can optionally specify an IP address that references an
	   existing static (reserved) IP address resource. When omitted, Google Cloud
	   assigns an ephemeral IP address.
	   Use one of the following formats to specify an IP address while creating a
	   forwarding rule:
	   - IP address number, as in `100.1.2.3`
	   - IPv6 address range, as in `2600:1234::/96`
	   - Full resource URL, as in
	   `https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name`
	   - Partial URL or by name, as in:
	   - `projects/project_id/regions/region/addresses/address-name`
	   - `regions/region/addresses/address-name`
	   - `global/addresses/address-name`
	   - `address-name`

	   The forwarding rule's `target`,
	   and in most cases, also the `loadBalancingScheme`, determine the
	   type of IP address that you can use. For detailed information, see
	   [IP address
	   specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	   When reading an `IPAddress`, the API always returns the IP
	   address number.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   This field identifies the subnetwork that the load balanced IP should
	   belong to for this Forwarding Rule, used in internal load balancing and
	   network load balancing with IPv6.
	   If the network specified is in auto subnet mode, this field is optional.
	   However, a subnetwork must be specified if the network is in custom subnet
	   mode or when creating external forwarding rule with IPv6.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specifies the forwarding rule type.
	   For more information about forwarding rules, refer to
	   [Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts).
	   Default value is `EXTERNAL`.
	   Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL_MANAGED`, `INTERNAL_SELF_MANAGED`.
	*/
	LoadBalancingScheme string `json:"loadBalancingScheme,omitempty" yaml:"loadBalancingScheme,omitempty"`

	/*
	   Opaque filter criteria used by Loadbalancer to restrict routing
	   configuration to a limited set xDS compliant clients. In their xDS
	   requests to Loadbalancer, xDS clients present node metadata. If a
	   match takes place, the relevant routing configuration is made available
	   to those proxies.
	   For each metadataFilter in this list, if its filterMatchCriteria is set
	   to MATCH_ANY, at least one of the filterLabels must match the
	   corresponding label provided in the metadata. If its filterMatchCriteria
	   is set to MATCH_ALL, then all of its filterLabels must match with
	   corresponding labels in the provided metadata.
	   metadataFilters specified here can be overridden by those specified in
	   the UrlMap that this ForwardingRule references.
	   metadataFilters only applies to Loadbalancers that have their
	   loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	   Structure is documented below.
	*/
	MetadataFilters []types.Compute_GlobalForwardingRuleMetadataFilter `json:"metadataFilters,omitempty" yaml:"metadataFilters,omitempty"`

	/*
	   This field is not used for external load balancing.
	   For Internal TCP/UDP Load Balancing, this field identifies the network that
	   the load balanced IP should belong to for this Forwarding Rule.
	   If the subnetwork is specified, the network of the subnetwork will be used.
	   If neither subnetwork nor this field is specified, the default network will
	   be used.
	   For Private Service Connect forwarding rules that forward traffic to Google
	   APIs, a network must be provided.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Service Directory resources to register this forwarding rule with.
	   Currently, only supports a single Service Directory resource.
	   Structure is documented below.
	*/
	ServiceDirectoryRegistrations types.Compute_GlobalForwardingRuleServiceDirectoryRegistrations `json:"serviceDirectoryRegistrations,omitempty" yaml:"serviceDirectoryRegistrations,omitempty"`

	/*
	   The URL of the target resource to receive the matched traffic.  For
	   regional forwarding rules, this target must be in the same region as the
	   forwarding rule. For global forwarding rules, this target must be a global
	   load balancing resource.
	   The forwarded traffic must be of a type appropriate to the target object.
	   -  For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	   -  For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle:
	   -  `vpc-sc` - [ APIs that support VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/supported-products).
	   -  `all-apis` - [All supported Google APIs](https://cloud.google.com/vpc/docs/private-service-connect#supported-apis).

	   For Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment.


	   - - -
	*/
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	/*
	   The IP Version that will be used by this global forwarding rule.
	   Possible values are: `IPV4`, `IPV6`.
	*/
	IpVersion string `json:"ipVersion,omitempty" yaml:"ipVersion,omitempty"`

	/*
	   Labels to apply to this forwarding rule.  A list of key->value pairs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	SourceIpRanges []string `json:"sourceIpRanges,omitempty" yaml:"sourceIpRanges,omitempty"`
}
