package types

type Ec2_getNetworkInsightsAnalysisReturnPathComponent struct {
	//
	SecurityGroupRules []Ec2_getNetworkInsightsAnalysisReturnPathComponentSecurityGroupRule `json:"securityGroupRules,omitempty" yaml:"securityGroupRules,omitempty"`

	//
	AclRules []Ec2_getNetworkInsightsAnalysisReturnPathComponentAclRule `json:"aclRules,omitempty" yaml:"aclRules,omitempty"`

	//
	SourceVpcs []Ec2_getNetworkInsightsAnalysisReturnPathComponentSourceVpc `json:"sourceVpcs,omitempty" yaml:"sourceVpcs,omitempty"`

	//
	Vpcs []Ec2_getNetworkInsightsAnalysisReturnPathComponentVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`

	//
	AdditionalDetails []Ec2_getNetworkInsightsAnalysisReturnPathComponentAdditionalDetail `json:"additionalDetails,omitempty" yaml:"additionalDetails,omitempty"`

	//
	AttachedTos []Ec2_getNetworkInsightsAnalysisReturnPathComponentAttachedTo `json:"attachedTos,omitempty" yaml:"attachedTos,omitempty"`

	//
	Components []Ec2_getNetworkInsightsAnalysisReturnPathComponentComponent `json:"components,omitempty" yaml:"components,omitempty"`

	//
	InboundHeaders []Ec2_getNetworkInsightsAnalysisReturnPathComponentInboundHeader `json:"inboundHeaders,omitempty" yaml:"inboundHeaders,omitempty"`

	//
	OutboundHeaders []Ec2_getNetworkInsightsAnalysisReturnPathComponentOutboundHeader `json:"outboundHeaders,omitempty" yaml:"outboundHeaders,omitempty"`

	//
	Subnets []Ec2_getNetworkInsightsAnalysisReturnPathComponentSubnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	//
	TransitGateways []Ec2_getNetworkInsightsAnalysisReturnPathComponentTransitGateway `json:"transitGateways,omitempty" yaml:"transitGateways,omitempty"`

	//
	DestinationVpcs []Ec2_getNetworkInsightsAnalysisReturnPathComponentDestinationVpc `json:"destinationVpcs,omitempty" yaml:"destinationVpcs,omitempty"`

	//
	RouteTableRoutes []Ec2_getNetworkInsightsAnalysisReturnPathComponentRouteTableRoute `json:"routeTableRoutes,omitempty" yaml:"routeTableRoutes,omitempty"`

	//
	SequenceNumber int `json:"sequenceNumber,omitempty" yaml:"sequenceNumber,omitempty"`

	//
	TransitGatewayRouteTableRoutes []Ec2_getNetworkInsightsAnalysisReturnPathComponentTransitGatewayRouteTableRoute `json:"transitGatewayRouteTableRoutes,omitempty" yaml:"transitGatewayRouteTableRoutes,omitempty"`
}
