package ec2clientvpn

type AuthorizationRule struct {
	// The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
	AccessGroupId string `json:"accessGroupId,omitempty" yaml:"accessGroupId,omitempty"`

	// Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
	AuthorizeAllGroups bool `json:"authorizeAllGroups,omitempty" yaml:"authorizeAllGroups,omitempty"`

	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId string `json:"clientVpnEndpointId,omitempty" yaml:"clientVpnEndpointId,omitempty"`

	// A brief description of the authorization rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
	TargetNetworkCidr string `json:"targetNetworkCidr,omitempty" yaml:"targetNetworkCidr,omitempty"`
}
