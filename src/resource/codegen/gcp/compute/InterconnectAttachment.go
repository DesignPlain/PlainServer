package compute

type InterconnectAttachment struct {
	/*
	   Provisioned bandwidth capacity for the interconnect attachment.
	   For attachments of type DEDICATED, the user can set the bandwidth.
	   For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	   Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	   Defaults to BPS_10G
	   Possible values are: `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, `BPS_50G`.
	*/
	Bandwidth string `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`

	/*
	   Up to 16 candidate prefixes that can be used to restrict the allocation
	   of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	   All prefixes must be within link-local address space (169.254.0.0/16)
	   and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	   an unused /29 from the supplied candidate prefix(es). The request will
	   fail if all possible /29s are in use on Google's edge. If not supplied,
	   Google will randomly select an unused /29 from all of link-local space.
	*/
	CandidateSubnets []string `json:"candidateSubnets,omitempty" yaml:"candidateSubnets,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Maximum Transmission Unit (MTU), in bytes, of packets passing through
	   this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
	*/
	Mtu string `json:"mtu,omitempty" yaml:"mtu,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is created. The
	   name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	   name must be 1-63 characters long and match the regular expression
	   `a-z?` which means the first character must be a
	   lowercase letter, and all following characters must be a dash, lowercase
	   letter, or digit, except the last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   URL of the cloud router to be used for dynamic routing. This router must be in
	   the same region as this InterconnectAttachment. The InterconnectAttachment will
	   automatically connect the Interconnect to the network & region within which the
	   Cloud Router is configured.
	*/
	Router string `json:"router,omitempty" yaml:"router,omitempty"`

	/*
	   The type of InterconnectAttachment you wish to create. Defaults to
	   DEDICATED.
	   Possible values are: `DEDICATED`, `PARTNER`, `PARTNER_PROVIDER`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   URL of the underlying Interconnect object that this attachment's
	   traffic will traverse through. Required if type is DEDICATED, must not
	   be set if type is PARTNER.
	*/
	Interconnect string `json:"interconnect,omitempty" yaml:"interconnect,omitempty"`

	// Region where the regional interconnect attachment resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Whether the VLAN attachment is enabled or disabled.  When using
	   PARTNER type this will Pre-Activate the interconnect attachment
	*/
	AdminEnabled bool `json:"adminEnabled,omitempty" yaml:"adminEnabled,omitempty"`

	/*
	   Desired availability domain for the attachment. Only available for type
	   PARTNER, at creation time. For improved reliability, customers should
	   configure a pair of attachments with one per availability domain. The
	   selected availability domain will be provided to the Partner via the
	   pairing key so that the provisioned circuit will lie in the specified
	   domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	*/
	EdgeAvailabilityDomain string `json:"edgeAvailabilityDomain,omitempty" yaml:"edgeAvailabilityDomain,omitempty"`

	/*
	   Indicates the user-supplied encryption option of this interconnect
	   attachment. Can only be specified at attachment creation for PARTNER or
	   DEDICATED attachments.
	   - NONE - This is the default value, which means that the VLAN attachment
	   carries unencrypted traffic. VMs are able to send traffic to, or receive
	   traffic from, such a VLAN attachment.
	   - IPSEC - The VLAN attachment carries only encrypted traffic that is
	   encrypted by an IPsec device, such as an HA VPN gateway or third-party
	   IPsec VPN. VMs cannot directly send traffic to, or receive traffic from,
	   such a VLAN attachment. To use HA VPN over Cloud Interconnect, the VLAN
	   attachment must be created with this option.
	   Default value is `NONE`.
	   Possible values are: `NONE`, `IPSEC`.
	*/
	Encryption string `json:"encryption,omitempty" yaml:"encryption,omitempty"`

	/*
	   URL of addresses that have been reserved for the interconnect attachment,
	   Used only for interconnect attachment that has the encryption option as
	   IPSEC.
	   The addresses must be RFC 1918 IP address ranges. When creating HA VPN
	   gateway over the interconnect attachment, if the attachment is configured
	   to use an RFC 1918 IP address, then the VPN gateway's IP address will be
	   allocated from the IP address range specified here.
	   For example, if the HA VPN gateway's interface 0 is paired to this
	   interconnect attachment, then an RFC 1918 IP address for the VPN gateway
	   interface 0 will be allocated from the IP address specified for this
	   interconnect attachment.
	   If this field is not specified for interconnect attachment that has
	   encryption option as IPSEC, later on when creating HA VPN gateway on this
	   interconnect attachment, the HA VPN gateway's IP address will be
	   allocated from regional external IP address pool.
	*/
	IpsecInternalAddresses []string `json:"ipsecInternalAddresses,omitempty" yaml:"ipsecInternalAddresses,omitempty"`

	/*
	   The stack type for this interconnect attachment to identify whether the IPv6
	   feature is enabled or not. If not specified, IPV4_ONLY will be used.
	   This field can be both set at interconnect attachments creation and update
	   interconnect attachment operations.
	   Possible values are: `IPV4_IPV6`, `IPV4_ONLY`.
	*/
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`

	/*
	   The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	   using PARTNER type this will be managed upstream.
	*/
	VlanTag8021q int `json:"vlanTag8021q,omitempty" yaml:"vlanTag8021q,omitempty"`
}
