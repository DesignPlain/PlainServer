package types

type Ec2_getNetworkInsightsAnalysisReturnPathComponentTransitGatewayRouteTableRoute struct {
	//
	PrefixListId string `json:"prefixListId,omitempty" yaml:"prefixListId,omitempty"`

	//
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	//
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	//
	RouteOrigin string `json:"routeOrigin,omitempty" yaml:"routeOrigin,omitempty"`

	//
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	//
	AttachmentId string `json:"attachmentId,omitempty" yaml:"attachmentId,omitempty"`

	//
	DestinationCidr string `json:"destinationCidr,omitempty" yaml:"destinationCidr,omitempty"`
}
