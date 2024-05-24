package types

type Networkmanager_getCoreNetworkPolicyDocumentSegmentAction struct {
	// A set subtraction of segments to not share with.
	ShareWithExcepts []string `json:"shareWithExcepts,omitempty" yaml:"shareWithExcepts,omitempty"`

	// A list of strings to share with. Must be a substring is all segments. Valid values include: `["-"]` or `["<segment-names>"]`.
	ShareWiths []string `json:"shareWiths,omitempty" yaml:"shareWiths,omitempty"`

	// Action to take for the chosen segment. Valid values `create-route` or `share`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// A user-defined string describing the segment action.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of strings containing CIDRs. You can define the IPv4 and IPv6 CIDR notation for each AWS Region. For example, `10.1.0.0/16` or `2001:db8::/56`. This is an array of CIDR notation strings.
	DestinationCidrBlocks []string `json:"destinationCidrBlocks,omitempty" yaml:"destinationCidrBlocks,omitempty"`

	// A list of strings. Valid values include `["blackhole"]` or a list of attachment ids.
	Destinations []string `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// String. This mode places the attachment and return routes in each of the `share_with` segments. Valid values include: `attachment-route`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// Name of the segment.
	Segment string `json:"segment,omitempty" yaml:"segment,omitempty"`
}
