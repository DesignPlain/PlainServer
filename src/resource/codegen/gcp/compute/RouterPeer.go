package compute

import types "libds/gcp/types"

type RouterPeer struct {
	/*
	   IPv6 address of the interface inside Google Cloud Platform.
	   The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	   If you do not specify the next hop addresses, Google Cloud automatically
	   assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	*/
	Ipv6NexthopAddress string `json:"ipv6NexthopAddress,omitempty" yaml:"ipv6NexthopAddress,omitempty"`

	/*
	   Peer BGP Autonomous System Number (ASN).
	   Each BGP interface may use a different value.
	*/
	PeerAsn int `json:"peerAsn,omitempty" yaml:"peerAsn,omitempty"`

	/*
	   IP address of the BGP interface outside Google Cloud Platform.
	   Only IPv4 is supported. Required if `ip_address` is set.
	*/
	PeerIpAddress string `json:"peerIpAddress,omitempty" yaml:"peerIpAddress,omitempty"`

	/*
	   The name of the Cloud Router in which this BgpPeer will be configured.


	   - - -
	*/
	Router string `json:"router,omitempty" yaml:"router,omitempty"`

	/*
	   User-specified flag to indicate which mode to use for advertisement.
	   Valid values of this enum field are: `DEFAULT`, `CUSTOM`
	   Default value is `DEFAULT`.
	   Possible values are: `DEFAULT`, `CUSTOM`.
	*/
	AdvertiseMode string `json:"advertiseMode,omitempty" yaml:"advertiseMode,omitempty"`

	/*
	   The priority of routes advertised to this BGP peer.
	   Where there is more than one matching route of maximum
	   length, the routes with the lowest priority value win.
	*/
	AdvertisedRoutePriority int `json:"advertisedRoutePriority,omitempty" yaml:"advertisedRoutePriority,omitempty"`

	// Name of the interface the BGP peer is associated with.
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`

	/*
	   Name of this BGP peer. The name must be 1-63 characters long,
	   and comply with RFC1035. Specifically, the name must be 1-63 characters
	   long and match the regular expression `a-z?` which
	   means the first character must be a lowercase letter, and all
	   following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Region where the router and BgpPeer reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   User-specified list of prefix groups to advertise in custom
	   mode, which currently supports the following option:
	*/
	AdvertisedGroups []string `json:"advertisedGroups,omitempty" yaml:"advertisedGroups,omitempty"`

	/*
	   The status of the BGP peer connection. If set to false, any active session
	   with the peer is terminated and all associated routing information is removed.
	   If set to true, the peer connection can be established with routing information.
	   The default is true.
	*/
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	/*
	   IP address of the interface inside Google Cloud Platform.
	   Only IPv4 is supported.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   IPv6 address of the BGP interface outside Google Cloud Platform.
	   The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	   If you do not specify the next hop addresses, Google Cloud automatically
	   assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	*/
	PeerIpv6NexthopAddress string `json:"peerIpv6NexthopAddress,omitempty" yaml:"peerIpv6NexthopAddress,omitempty"`

	/*
	   The URI of the VM instance that is used as third-party router appliances
	   such as Next Gen Firewalls, Virtual Routers, or Router Appliances.
	   The VM instance must be located in zones contained in the same region as
	   this Cloud Router. The VM instance is the peer side of the BGP session.
	*/
	RouterApplianceInstance string `json:"routerApplianceInstance,omitempty" yaml:"routerApplianceInstance,omitempty"`

	/*
	   User-specified list of individual IP ranges to advertise in
	   custom mode. This field can only be populated if advertiseMode
	   is `CUSTOM` and is advertised to all peers of the router. These IP
	   ranges will be advertised in addition to any specified groups.
	   Leave this field blank to advertise no custom IP ranges.
	   Structure is documented below.
	*/
	AdvertisedIpRanges []types.Compute_RouterPeerAdvertisedIpRange `json:"advertisedIpRanges,omitempty" yaml:"advertisedIpRanges,omitempty"`

	/*
	   BFD configuration for the BGP peering.
	   Structure is documented below.
	*/
	Bfd types.Compute_RouterPeerBfd `json:"bfd,omitempty" yaml:"bfd,omitempty"`

	// Enable IPv6 traffic over BGP Peer. If not specified, it is disabled by default.
	EnableIpv6 bool `json:"enableIpv6,omitempty" yaml:"enableIpv6,omitempty"`

	/*
	   Present if MD5 authentication is enabled for the peering. Must be the name of one of the entries in the
	   Router.md5_authentication_keys. The field must comply with RFC1035.
	*/
	Md5AuthenticationKey types.Compute_RouterPeerMd5AuthenticationKey `json:"md5AuthenticationKey,omitempty" yaml:"md5AuthenticationKey,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
