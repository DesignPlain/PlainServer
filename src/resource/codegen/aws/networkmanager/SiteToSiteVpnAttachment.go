package networkmanager

type SiteToSiteVpnAttachment struct {
	// The ID of a core network for the VPN attachment.
	CoreNetworkId string `json:"coreNetworkId,omitempty" yaml:"coreNetworkId,omitempty"`

	// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The ARN of the site-to-site VPN connection.

	   The following arguments are optional:
	*/
	VpnConnectionArn string `json:"vpnConnectionArn,omitempty" yaml:"vpnConnectionArn,omitempty"`
}
