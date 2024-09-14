package types

type Networkmanager_getCoreNetworkPolicyDocumentSegmentAction struct {
	// Name of the segment.
	Segment string `json:"segment,omitempty" yaml:"segment,omitempty"`

	// A list of strings to share with. Must be a substring is all segments. Valid values include: `["-"]` or `["<segment-names>"]`.
	ShareWiths []string `json:"shareWiths,omitempty" yaml:"shareWiths,omitempty"`

	// The network function groups and any edge overrides associated with the action.
	Via Networkmanager_getCoreNetworkPolicyDocumentSegmentActionVia `json:"via,omitempty" yaml:"via,omitempty"`

	// The destination segments for the `send-via` or `send-to` `action`.
	WhenSentTo Networkmanager_getCoreNetworkPolicyDocumentSegmentActionWhenSentTo `json:"whenSentTo,omitempty" yaml:"whenSentTo,omitempty"`

	// Action to take for the chosen segment. Valid values: `create-route`, `share`, `send-via` and `send-to`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// A user-defined string describing the segment action.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of strings containing CIDRs. You can define the IPv4 and IPv6 CIDR notation for each AWS Region. For example, `10.1.0.0/16` or `2001:db8::/56`. This is an array of CIDR notation strings.
	DestinationCidrBlocks []string `json:"destinationCidrBlocks,omitempty" yaml:"destinationCidrBlocks,omitempty"`

	// A list of strings. Valid values include `["blackhole"]` or a list of attachment ids.
	Destinations []string `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// String. When `action` is `share`, a `mode` value of `attachment-route` places the attachment and return routes in each of the `share_with` segments. When `action` is `send-via`, indicates the mode used for packets. Valid values: `attachment-route`, `single-hop`, `dual-hop`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// A set subtraction of segments to not share with.
	ShareWithExcepts []string `json:"shareWithExcepts,omitempty" yaml:"shareWithExcepts,omitempty"`
}
