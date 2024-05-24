package ec2transitgateway

type TransitGateway struct {
	/*
	   Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.

	   > --NOTE:-- Modifying `amazon_side_asn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazon_side_asn`.
	*/
	AmazonSideAsn int `json:"amazonSideAsn,omitempty" yaml:"amazonSideAsn,omitempty"`

	// Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTableAssociation string `json:"defaultRouteTableAssociation,omitempty" yaml:"defaultRouteTableAssociation,omitempty"`

	// Description of the EC2 Transit Gateway.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether Multicast support is enabled. Required to use `ec2_transit_gateway_multicast_domain`. Valid values: `disable`, `enable`. Default value: `disable`.
	MulticastSupport string `json:"multicastSupport,omitempty" yaml:"multicastSupport,omitempty"`

	// Key-value tags for the EC2 Transit Gateway. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	VpnEcmpSupport string `json:"vpnEcmpSupport,omitempty" yaml:"vpnEcmpSupport,omitempty"`

	// Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
	AutoAcceptSharedAttachments string `json:"autoAcceptSharedAttachments,omitempty" yaml:"autoAcceptSharedAttachments,omitempty"`

	// Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTablePropagation string `json:"defaultRouteTablePropagation,omitempty" yaml:"defaultRouteTablePropagation,omitempty"`

	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport string `json:"dnsSupport,omitempty" yaml:"dnsSupport,omitempty"`

	// One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
	TransitGatewayCidrBlocks []string `json:"transitGatewayCidrBlocks,omitempty" yaml:"transitGatewayCidrBlocks,omitempty"`
}
