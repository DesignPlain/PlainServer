package ec2

type Eip struct {
	// Network interface ID to associate with.
	NetworkInterface string `json:"networkInterface,omitempty" yaml:"networkInterface,omitempty"`

	/*
	   EC2 IPv4 address pool identifier or `amazon`.
	   This option is only available for VPC EIPs.
	*/
	PublicIpv4Pool string `json:"publicIpv4Pool,omitempty" yaml:"publicIpv4Pool,omitempty"`

	// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIp string `json:"associateWithPrivateIp,omitempty" yaml:"associateWithPrivateIp,omitempty"`

	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
	CustomerOwnedIpv4Pool string `json:"customerOwnedIpv4Pool,omitempty" yaml:"customerOwnedIpv4Pool,omitempty"`

	// Indicates if this EIP is for use in VPC (`vpc`).
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// EC2 instance ID.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup string `json:"networkBorderGroup,omitempty" yaml:"networkBorderGroup,omitempty"`

	/*
	   Boolean if the EIP is in a VPC or not. Use `domain` instead.
	   Defaults to `true` unless the region supports EC2-Classic.

	   > --NOTE:-- You can specify either the `instance` ID or the `network_interface` ID, but not both. Including both will --not-- return an error from the AWS API, but will have undefined behavior. See the relevant [AssociateAddress API Call][1] for more information.

	   > --NOTE:-- Specifying both `public_ipv4_pool` and `address` won't cause an error but `address` will be used in the
	   case both options are defined as the api only requires one or the other.
	*/
	Vpc bool `json:"vpc,omitempty" yaml:"vpc,omitempty"`

	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
}
