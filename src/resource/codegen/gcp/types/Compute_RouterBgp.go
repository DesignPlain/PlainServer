package types

type Compute_RouterBgp struct {
	/*
	   Local BGP Autonomous System Number (ASN). Must be an RFC6996
	   private ASN, either 16-bit or 32-bit. The value will be fixed for
	   this router resource. All VPN tunnels that link to this router
	   will have the same local ASN.
	*/
	Asn int `json:"asn,omitempty" yaml:"asn,omitempty"`

	/*
	   The interval in seconds between BGP keepalive messages that are sent
	   to the peer. Hold time is three times the interval at which keepalive
	   messages are sent, and the hold time is the maximum number of seconds
	   allowed to elapse between successive keepalive messages that BGP
	   receives from a peer.
	   BGP will use the smaller of either the local hold time value or the
	   peer's hold time value as the hold time for the BGP connection
	   between the two peers. If set, this value must be between 20 and 60.
	   The default is 20.
	*/
	KeepaliveInterval int `json:"keepaliveInterval,omitempty" yaml:"keepaliveInterval,omitempty"`

	/*
	   User-specified flag to indicate which mode to use for advertisement.
	   Default value is `DEFAULT`.
	   Possible values are: `DEFAULT`, `CUSTOM`.
	*/
	AdvertiseMode string `json:"advertiseMode,omitempty" yaml:"advertiseMode,omitempty"`

	/*
	   User-specified list of prefix groups to advertise in custom mode.
	   This field can only be populated if advertiseMode is CUSTOM and
	   is advertised to all peers of the router. These groups will be
	   advertised in addition to any specified prefixes. Leave this field
	   blank to advertise no custom groups.
	   This enum field has the one valid value: ALL_SUBNETS
	*/
	AdvertisedGroups []string `json:"advertisedGroups,omitempty" yaml:"advertisedGroups,omitempty"`

	/*
	   User-specified list of individual IP ranges to advertise in
	   custom mode. This field can only be populated if advertiseMode
	   is CUSTOM and is advertised to all peers of the router. These IP
	   ranges will be advertised in addition to any specified groups.
	   Leave this field blank to advertise no custom IP ranges.
	   Structure is documented below.
	*/
	AdvertisedIpRanges []Compute_RouterBgpAdvertisedIpRange `json:"advertisedIpRanges,omitempty" yaml:"advertisedIpRanges,omitempty"`
}
