package types

type Ec2transitgateway_getRouteTableRoutesRoute struct {
	// The CIDR used for route destination matches.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`

	// The ID of the prefix list used for destination matches.
	PrefixListId string `json:"prefixListId,omitempty" yaml:"prefixListId,omitempty"`

	// The current state of the route, can be `active`, `deleted`, `pending`, `blackhole`, `deleting`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// The id of the transit gateway route table announcement, most of the time it is an empty string.
	TransitGatewayRouteTableAnnouncementId string `json:"transitGatewayRouteTableAnnouncementId,omitempty" yaml:"transitGatewayRouteTableAnnouncementId,omitempty"`

	// The type of the route, can be `propagated` or `static`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
